package core

import (
	"errors"
)

var (
	ErrLinkNotFound = errors.New("link not found")
	ErrNotFoundDocs = errors.New("no documents found")
)
