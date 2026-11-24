package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/sikozonpc/ecom/internal/env"
)

func main() {
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dns: env.GetString("GOOSE_DBSTRING", ""),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := pgx.Connect(ctx, cfg.db.dns)
	if err != nil {
		os.Exit(1)
	}

	defer conn.Close(ctx)

	logger.Info("Database has been connected")

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server fail to start", "error", err)
		os.Exit(1)
	}
}
