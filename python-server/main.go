package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type ExecuteRequest struct {
	Code string `json:"code"`
}

type ExecuteResponse struct {
	Status  string `json:"status"`
	Output  string `json:"output,omitempty"`
	Message string `json:"message,omitempty"`
}

func main() {
	http.HandleFunc("/execute", handleExecute)

	port := "9010"
	log.Printf("Python execution server is running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}
}

// JSON形式のエラーレスポンスを返すためのヘルパー関数
func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ExecuteResponse{
		Status:  "error",
		Message: message,
	})
}

func handleExecute(w http.ResponseWriter, r *http.Request) {
	allowedReferer := os.Getenv("ALLOWED_REFERER")
	// Check the Referer header
	referer := r.Header.Get("Referer")
	if referer != allowedReferer {
		log.Printf("Invalid Referer: %s\n", referer)
		writeJSONError(w, "Forbidden: Access denied", http.StatusForbidden)
		return
	}

	if r.Method != http.MethodPost {
		writeJSONError(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req ExecuteRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSONError(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &req); err != nil {
		writeJSONError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Write Python code to a temporary file
	tempDir, err := os.MkdirTemp("", "python-server")
	if err != nil {
		writeJSONError(w, "Failed to create temporary directory", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tempDir)

	scriptPath := filepath.Join(tempDir, "script.py")
	if err := os.WriteFile(scriptPath, []byte(req.Code), 0644); err != nil {
		writeJSONError(w, "Failed to write script file", http.StatusInternalServerError)
		return
	}

	// Execute the Python code with a 2-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	execCmd := exec.CommandContext(ctx, "python3", scriptPath)
	var execOut, execErr bytes.Buffer
	execCmd.Stdout = &execOut
	execCmd.Stderr = &execErr
	if err := execCmd.Run(); err != nil {
		// Check if the error is a timeout
		if ctx.Err() == context.DeadlineExceeded {
			resp := ExecuteResponse{
				Status:  "error",
				Message: "Execution timed out",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
		// Handle normal execution errors
		resp := ExecuteResponse{
			Status:  "error",
			Message: execErr.String(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	log.Println("Output:", execOut.String())

	// Send the output back to the client
	resp := ExecuteResponse{
		Status: "success",
		Output: execOut.String(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
