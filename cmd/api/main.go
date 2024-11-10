package main

import (
	"log"

	"github.com/mosmartin/go-social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetStringEnv("ADDR", ":8080"),
	}

	app := &app{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
