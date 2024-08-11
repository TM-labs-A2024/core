package controller

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/TM-labs-A2024/core/services/backend-server/pkg/config"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/storage/dropbox"
	"github.com/TM-labs-A2024/core/services/backend-server/pkg/blockchain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageClient interface {
	UploadFile(uuid.UUID, io.Reader) (string, error)
	GenerateURL(string) (*url.URL, error)
	DeleteFile(string) error
}

type Controller struct {
	pool            *pgxpool.Pool
	queries         *db.Queries
	storage         StorageClient
	logger          *slog.Logger
	blockchain      *blockchain.Client
	ivEncryptionKey string
}

func NewController(conf config.Config, logger *slog.Logger) (*Controller, error) {
	ctx := context.Background()

	storage, err := dropbox.New(conf.StorageAPIKey)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.New(ctx, conf.DatabaseURL)
	if err != nil {
		return nil, err
	}

	blockchain, err := blockchain.New(
		conf.ChaincodeName,
		conf.ChannelName,
		conf.CryptoPath,
		conf.PeerEndpoint,
	)
	if err != nil {
		return nil, err
	}

	queries := db.New(pool)

	controller := &Controller{
		queries:         queries,
		logger:          logger,
		pool:            pool,
		storage:         storage,
		blockchain:      blockchain,
		ivEncryptionKey: conf.IVEncryptionKey,
	}

	return controller, nil
}
