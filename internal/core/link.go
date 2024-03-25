package core

import "fmt"

type Link struct {
	UUID         string `bson:"_id" json:"uuid"`
	OriginalLink string `bson:"original_link" json:"original_link"`
	ShortLink    string `bson:"short_link" json:"short_link"`
	Count        int    `bson:"count" json:"count"`
	CreatedAt    string `bson:"created_at" json:"created_at"`
}

type CreateLinkRequest struct {
	OriginalLink string `json:"original_link"`
}

func NewShortLink(uuid string) string {
	shorUri := fmt.Sprintf("http://localhost:3000/%s", uuid)

	return shorUri
}

type DataResponse struct {
	CreatedAt string `json:"created_at"`
	Count     int    `json:"count"`
}
