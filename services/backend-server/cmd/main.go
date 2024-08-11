package main

import (
	"fmt"
	"log/slog"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/server"
	"github.com/TM-labs-A2024/core/services/backend-server/pkg/config"
)

func main() {
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic(fmt.Errorf("could not load config %w", err))
	}

	s, err := server.NewServer(config)
	if err != nil {
		slog.Error("could not create server", slog.Any("error", err))
	}

	s.Logger.Error("faltal error while running server", slog.Any("error", s.Start(":8080")))
}
