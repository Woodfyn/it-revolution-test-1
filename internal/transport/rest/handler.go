package rest

import (
	"net/http"

	"github.com/Woodfyn/it-revolution-test-1/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	{
		api.HandleFunc("/transform", h.TransformLink).Methods(http.MethodPost)
		api.HandleFunc("/original/{uuid}", h.OriginalLink).Methods(http.MethodGet)
		api.HandleFunc("/statistics", h.GetStatistics).Methods(http.MethodGet)
		api.HandleFunc("/statistics/{uuid}", h.GetStatisticsById).Methods(http.MethodGet)
	}

	return r
}
