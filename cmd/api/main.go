package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &app{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}