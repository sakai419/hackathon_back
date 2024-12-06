package main

import (
	"bytes"
	"fmt"
	v1 "local-test/api/v1"
	"local-test/internal/config"
	"os"
	"os/exec"
	"time"

	// "local-test/internal/model"
	"local-test/pkg/database"
	"local-test/pkg/firebase"
	"local-test/pkg/utils"
	"log"
)

func main() {
    // サンプルのC言語コード
    cCode := `
    #include <stdio.h>
    int main() {
        printf("Hello, World!\n");
        return 0;
    }
    `

    // Cコードを一時ファイルに保存
    tmpFile, err := os.CreateTemp("", "*.c")
    if err != nil {
        panic(err)
    }
    defer os.Remove(tmpFile.Name()) // ファイルをクリーンアップ
    defer tmpFile.Close()

    if _, err := tmpFile.Write([]byte(cCode)); err != nil {
        panic(err)
    }

    // GCCでコンパイル
    execFile := tmpFile.Name() + ".out"
    cmd := exec.Command("gcc", tmpFile.Name(), "-o", execFile)
    if err := cmd.Run(); err != nil {
        panic(err)
    }
    defer os.Remove(execFile) // 実行ファイルをクリーンアップ

    // 実行可能ファイルを実行
    execCmd := exec.Command(execFile)
    var outBuf, errBuf bytes.Buffer
    execCmd.Stdout = &outBuf
    execCmd.Stderr = &errBuf

    // タイムアウト付きで実行
    err = execCmd.Start()
    if err != nil {
        panic(err)
    }
    done := make(chan error, 1)
    go func() { done <- execCmd.Wait() }()

    select {
    case <-time.After(2 * time.Second): // 2秒タイムアウト
        execCmd.Process.Kill()
        fmt.Println("Execution timed out")
    case err := <-done:
        if err != nil {
            fmt.Printf("Error: %s\n", errBuf.String())
        } else {
            fmt.Printf("Output: %s\n", outBuf.String())
        }
    }

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Connect to database
	db, err := database.ConnectToDB(cfg.DBConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	firebaseClient, err := firebase.InitFirebaseClient(cfg.FirebaseConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Close DB connection on syscall
	utils.CloseDBWithSysCall(db)

	// Setup server
	server := v1.NewServer(db, firebaseClient, cfg.ServerConfig)

	// Start server
	if err := server.Start(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}