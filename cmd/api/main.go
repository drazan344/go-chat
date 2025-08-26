package main

import (
	"log"
	"github.com/drazan344/go-chat/internal/env"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or couldn't be loaded: %v", err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}


	mux := app.mount()
	log.Fatal(app.run(mux))
}