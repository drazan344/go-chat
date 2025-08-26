package main

import (
	"log"

	"gihub.com/drazan344/go-chat/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}


	mux := app.mount()
	log.Fatal(app.run(mux))
}