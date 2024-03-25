package core

import (
	"errors"
)

var (
	ErrUuidNotFound = errors.New("uuid not found")
	ErrLinkNotFound = errors.New("link not found")
	ErrNotFoundDocs = errors.New("no documents found")
)
