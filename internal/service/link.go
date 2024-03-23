package service

import "github.com/Woodfyn/it-revolution-test-1/internal/repository/mongo"

type Link struct {
	repo mongo.Linker
}

func NewLink(repo mongo.Linker) *Link {
	return &Link{
		repo: repo,
	}
}
