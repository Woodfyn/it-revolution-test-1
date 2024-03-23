package core

type Link struct {
	UUID         string
	OriginalLink string
	ShortLink    string
	Count        int
}

type CreateLinkRequest struct {
	OriginalLink string
}
