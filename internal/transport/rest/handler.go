package rest

import (
	"github.com/Woodfyn/it-revolution-test-1/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
