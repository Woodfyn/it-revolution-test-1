package mongo

import "go.mongodb.org/mongo-driver/mongo"

type Linker interface {
}

type Repository struct {
	Linker Linker
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Linker: NewLink(db),
	}
}
