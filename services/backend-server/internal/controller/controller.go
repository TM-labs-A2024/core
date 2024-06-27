package controller

import (
	"context"
	"log/slog"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/jackc/pgx/v5"
)

type Controller struct {
	conn    *pgx.Conn
	queries *db.Queries
	logger  *slog.Logger
}

func NewController(dbUrl string, logger *slog.Logger) (*Controller, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		return nil, err
	}

	queries := db.New(conn)

	controller := &Controller{
		queries: queries,
		logger:  logger,
		conn:    conn,
	}

	return controller, nil
}
