package service

import (
	"context"

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
	}

	shortLinkCheck, err := l.repo.GetByOriginalLink(ctx, originalLink)
	if err == nil {
		return shortLinkCheck, nil
	}

	shortLink, err := l.repo.AddLink(ctx, link)

	return shortLink, err
}

func (l *Link) OriginalLink(ctx context.Context, uuid string) (string, error) {
	return l.repo.GetByUUID(ctx, uuid)
}

func (l *Link) GetStatistics(ctx context.Context) ([]core.DataResponse, error) {
	return l.repo.GetStatistics(ctx)
}
