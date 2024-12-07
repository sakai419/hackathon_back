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

// JSON形式のエラーレスポンスを返すためのヘルパー関数
func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(CompileResponse{
		Status:  "error",
		Message: message,
	})
}

func handleCompile(w http.ResponseWriter, r *http.Request) {
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
	var req CompileRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSONError(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &req); err != nil {
		writeJSONError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Write C code to a temporary file
	tempDir, err := os.MkdirTemp("", "gcc-server")
	if err != nil {
		writeJSONError(w, "Failed to create temporary directory", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tempDir)

	sourcePath := filepath.Join(tempDir, "code.c")
	if err := os.WriteFile(sourcePath, []byte(req.Code), 0644); err != nil {
		writeJSONError(w, "Failed to write source file", http.StatusInternalServerError)
		return
	}

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

	// コードの実行に2秒のタイムアウトを設定
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	runCmd := exec.CommandContext(ctx, outputPath)
	var runOut, runErr bytes.Buffer
	runCmd.Stdout = &runOut
	runCmd.Stderr = &runErr
	if err := runCmd.Run(); err != nil {
		// タイムアウトかどうかをチェック
		if ctx.Err() == context.DeadlineExceeded {
			resp := CompileResponse{
				Status:  "error",
				Message: "Execution timed out",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
		// 通常のエラー
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
