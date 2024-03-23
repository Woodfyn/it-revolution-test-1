package mongo

import "go.mongodb.org/mongo-driver/mongo"

type Link struct {
	db *mongo.Database
}

func NewLink(db *mongo.Database) *Link {
	return &Link{
		db: db,
	}
}
