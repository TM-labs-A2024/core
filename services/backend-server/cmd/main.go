package main

import (
	"log/slog"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		slog.Error("could not create server", err)
	}

	s.Logger.Error("faltal error while running server", s.Start(":8080"))
}
