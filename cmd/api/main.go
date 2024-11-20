package main

import (
	"log"
	"log/slog"

	"github.com/mosmartin/go-social/internal/db"
	"github.com/mosmartin/go-social/internal/env"
	"github.com/mosmartin/go-social/internal/store"
)

const version = "0.0.1"

func main() {
	cfg := config{
		addr: env.GetStringEnv("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetStringEnv("DB_ADDR", "postgresql://localhost:5432"),
			maxOpenConns: env.GetIntEnv("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetIntEnv("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetStringEnv("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetStringEnv("ENV", "development"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	slog.Info("database connection established")

	dbStore := store.NewStorage(db)

	app := &app{
		config:  cfg,
		dbStore: dbStore,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
