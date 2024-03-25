package service

import (
	"context"

	"github.com/Woodfyn/it-revolution-test-1/internal/core"
	"github.com/Woodfyn/it-revolution-test-1/internal/repository/mongo"
)

type Linker interface {
	TransformLink(ctx context.Context, originalLink string) (string, error)
	OriginalLink(ctx context.Context, shortLink string) (string, error)
	GetAllStatistics(ctx context.Context) ([]core.DataResponse, error)
	GetStatisticsById(ctx context.Context, uuid string) (core.DataResponse, error)
}

type Service struct {
	Linker Linker
}

type Deps struct {
	MongoRepo *mongo.Repository
}

func NewService(deps Deps) *Service {
	return &Service{
		Linker: NewLink(deps.MongoRepo.Linker),
	}
}
