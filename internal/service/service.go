package service

import "github.com/Woodfyn/it-revolution-test-1/internal/repository/mongo"

type Linker interface {
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
