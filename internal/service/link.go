package service

import (
	"context"
	"time"

	"github.com/Woodfyn/it-revolution-test-1/internal/core"
	"github.com/Woodfyn/it-revolution-test-1/internal/repository/mongo"
	"github.com/google/uuid"
)

type Link struct {
	repo mongo.Linker
}

func NewLink(repo mongo.Linker) *Link {
	return &Link{
		repo: repo,
	}
}

func (l *Link) TransformLink(ctx context.Context, originalLink string) (string, error) {

	uuid := uuid.NewString()
	shortUuid := uuid[:4]

	link := core.Link{
		UUID:         shortUuid,
		OriginalLink: originalLink,
		ShortLink:    core.NewShortLink(shortUuid),
		Count:        1,
		CreatedAt:    time.Now().Add(time.Hour * 2).Format(time.DateTime),
	}

	shortLink, err := l.repo.AddLink(ctx, link)

	return shortLink, err
}

func (l *Link) OriginalLink(ctx context.Context, uuid string) (string, error) {
	return l.repo.GetByUUID(ctx, uuid)
}

func (l *Link) GetAllStatistics(ctx context.Context) ([]core.DataResponse, error) {
	return l.repo.GetAllStatistics(ctx)
}

func (l *Link) GetStatisticsById(ctx context.Context, uuid string) (core.DataResponse, error) {
	return l.repo.GetStatisticsById(ctx, uuid)
}
