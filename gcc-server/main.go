package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type CompileRequest struct {
	Code string `json:"code"`
}

type CompileResponse struct {
	Status  string `json:"status"`
	Output  string `json:"output,omitempty"`
	Message string `json:"message,omitempty"`
}

func main() {
	http.HandleFunc("/compile", handleCompile)

	port := "9000"
	log.Printf("GCC server is running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}
}

func handleCompile(w http.ResponseWriter, r *http.Request) {
	allowedReferer := os.Getenv("ALLOWED_REFERER")
	log.Printf("Allowed Referer: %s\n", allowedReferer)
	// Check the Referer header
	referer := r.Header.Get("Referer")
	if referer != allowedReferer {
		http.Error(w, "Forbidden: Access denied", http.StatusForbidden)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Request from: %s\n", r.RemoteAddr)

	// Parse request body
	var req CompileRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	log.Printf("Received code: %s\n", req.Code)

	// Write C code to a temporary file
	tempDir, err := os.MkdirTemp("", "gcc-server")
	if err != nil {
		http.Error(w, "Failed to create temporary directory", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tempDir)

	log.Printf("Temporary directory: %s\n", tempDir)

	sourcePath := filepath.Join(tempDir, "code.c")
	if err := os.WriteFile(sourcePath, []byte(req.Code), 0644); err != nil {
		http.Error(w, "Failed to write source file", http.StatusInternalServerError)
		return
	}

	log.Printf("Source file: %s\n", sourcePath)

	// Compile the C code using GCC
	outputPath := filepath.Join(tempDir, "output")
	compileCmd := exec.Command("gcc", sourcePath, "-o", outputPath)
	var compileErr bytes.Buffer
	compileCmd.Stderr = &compileErr
	if err := compileCmd.Run(); err != nil {
		resp := CompileResponse{
			Status:  "error",
			Message: compileErr.String(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Execute the compiled binary
	runCmd := exec.Command(outputPath)
	var runOut, runErr bytes.Buffer
	runCmd.Stdout = &runOut
	runCmd.Stderr = &runErr
	if err := runCmd.Run(); err != nil {
		resp := CompileResponse{
			Status:  "error",
			Message: runErr.String(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	log.Println("Output:", runOut.String())

	// Send the output back to the client
	resp := CompileResponse{
		Status: "success",
		Output: runOut.String(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
