package controller

import (
	"context"
	"log/slog"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Controller struct {
	pool    *pgxpool.Pool
	queries *db.Queries
	logger  *slog.Logger
}

func NewController(dbUrl string, logger *slog.Logger) (*Controller, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		return nil, err
	}

	queries := db.New(pool)

	controller := &Controller{
		queries: queries,
		logger:  logger,
		pool:    pool,
	}

	return controller, nil
}
