package main

import (
	"log/slog"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		slog.Error("could not create server", slog.Any("error", err))
	}

	s.Logger.Error("faltal error while running server", slog.Any("error", s.Start(":8080")))
}
