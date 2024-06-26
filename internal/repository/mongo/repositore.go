package mongo

import (
	"context"

	"github.com/Woodfyn/it-revolution-test-1/internal/core"
	"go.mongodb.org/mongo-driver/mongo"
)

type Linker interface {
	AddLink(ctx context.Context, link core.Link) (string, error)
	// GetByOriginalLink(ctx context.Context, originalLink string) (string, error)
	GetByUUID(ctx context.Context, uuid string) (string, error)
	GetAllStatistics(ctx context.Context) ([]core.DataResponse, error)
	GetStatisticsById(ctx context.Context, uuid string) (core.DataResponse, error)
}

type Repository struct {
	Linker Linker
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Linker: NewLink(db),
	}
}
