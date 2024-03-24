package core

import "fmt"

type Link struct {
	UUID         string `bson:"_id" json:"uuid"`
	OriginalLink string `bson:"original_link" json:"original_link"`
	ShortLink    string `bson:"short_link" json:"short_link"`
	Count        int    `bson:"count" json:"count"`
}

type CreateLinkRequest struct {
	OriginalLink string `json:"original_link"`
}

func NewShortLink(uuid string) string {
	shorUri := fmt.Sprintf("http://localhost:8080/%s", uuid)

	return shorUri
}
