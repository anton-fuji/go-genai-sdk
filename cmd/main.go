package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/anton-fuji/go-genai-sdk/internal/chat"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	model := flag.String("model", "gemini-2.0-flash", "model name")
	temp := flag.Float64("temp", 0.5, "temperature 0.0-1.0")
	flag.Parse()

	_ = godotenv.Load()
	// wd, _ := os.Getwd()
	// log.Println("CWD=", wd)
	// log.Println("API_KEY =", os.Getenv("GOOGLE_API_KEY"))
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_API_KEYが未設定です")
	}

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ch, err := chat.New(ctx, client, *model, float32(*temp))
	if err != nil {
		log.Fatal(err)
	}

	if err := ch.Run(ctx, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
