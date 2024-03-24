package mongo

import (
	"context"

	"github.com/Woodfyn/it-revolution-test-1/internal/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Link struct {
	db *mongo.Collection
}

func NewLink(db *mongo.Database) *Link {
	return &Link{
		db: db.Collection(LinksCollection),
	}
}

func (l *Link) AddLink(ctx context.Context, link core.Link) (string, error) {
	l.db.InsertOne(ctx, link)

	return link.ShortLink, nil
}

func (l *Link) GetByOriginalLink(ctx context.Context, originalLink string) (string, error) {
	var link core.Link

	res := l.db.FindOne(ctx, bson.M{"original_link": originalLink})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return "", core.ErrNotFoundDocs
		}
		return "", res.Err()
	}

	if err := res.Decode(&link); err != nil {
		return "", err
	}

	return link.ShortLink, nil
}

func (l *Link) GetByUUID(ctx context.Context, uuid string) (string, error) {
	var link core.Link

	res := l.db.FindOne(ctx, bson.M{"_id": uuid})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return "", core.ErrNotFoundDocs
		}
		return "", res.Err()
	}

	if err := res.Decode(&link); err != nil {
		return "", err
	}

	link.Count++

	_, err := l.db.UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": bson.M{"count": link.Count}})
	if err != nil {
		return "", err
	}

	return link.OriginalLink, nil
}

func (l *Link) GetStatistics(ctx context.Context, uuid string) (int, error) {
	var link core.Link

	res := l.db.FindOne(ctx, bson.M{"_id": uuid})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return 0, core.ErrNotFoundDocs
		}
		return 0, res.Err()
	}

	if err := res.Decode(&link); err != nil {
		return 0, err
	}

	return link.Count, nil
}
