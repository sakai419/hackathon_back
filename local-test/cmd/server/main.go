package main

import (
	"context"
	"fmt"
	v1 "local-test/api/v1"
	"local-test/internal/config"
	"local-test/pkg/database"
	"local-test/pkg/firebase"
	"local-test/pkg/utils"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
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
	server := v1.NewServer(db, firebaseClient)

    // Connect to generative AI
    ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.GeminiConfig.APIKey))
    if err != nil {
        log.Fatalf("Failed to create Gemini client: %v", err)
    }
    defer client.Close()

    model := client.GenerativeModel("gemini-pro")
    prompt := genai.Text("あなたの名前は何ですか？")
    resp, err := model.GenerateContent(ctx, prompt)
    if err != nil {
        log.Fatalf("Failed to generate content: %v", err)
    }

    for _, c := range resp.Candidates {
        for _, p := range c.Content.Parts {
            fmt.Println(p)
        }
    }

	// Start server
	if err := server.Start(cfg.ServerConfig.Port); err != nil {
		log.Fatalf("Error: %v", err)
	}


}

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/google/generative-ai-go/genai"
// 	"google.golang.org/api/option"
// )

// func main() {
//     ctx := context.Background()

// 	print(os.Getenv("GEMINI_API_KEY"))

//     client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
//     if err != nil {
//         log.Fatalf("Failed to create Gemini client: %v", err)
//     }
//     defer client.Close()

//     model := client.GenerativeModel("gemini-pro")
//     prompt := genai.Text("こんにちは、世界！")
//     resp, err := model.GenerateContent(ctx, prompt)
//     if err != nil {
//         log.Fatalf("Failed to generate content: %v", err)
//     }

//     for _, c := range resp.Candidates {
//         for _, p := range c.Content.Parts {
//             fmt.Println(p)
//         }
//     }
// }
