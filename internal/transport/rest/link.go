package rest

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/Woodfyn/it-revolution-test-1/internal/core"
	"github.com/gorilla/mux"
)

func (h *Handler) TransformLink(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Info("TransformLink", "err read body", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var req core.CreateLinkRequest
	if err = json.Unmarshal(reqBytes, &req); err != nil {
		slog.Info("TransformLink", "err unmarshal body", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	shortLink, err := h.service.Linker.TransformLink(r.Context(), req.OriginalLink)
	if err != nil {
		slog.Info("TransformLink", "err service call", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(shortLink)
	if err != nil {
		slog.Info("TransformLink", "err write body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) OriginalLink(w http.ResponseWriter, r *http.Request) {
	uuid, err := getIdFromRequest(r)
	if err != nil {
		slog.Info("OriginalLink", "err get shortLink", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	originalLink, err := h.service.Linker.OriginalLink(r.Context(), uuid)
	if err != nil {
		slog.Info("OriginalLink", "err service call", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(originalLink)
	if err != nil {
		slog.Info("OriginalLink", "err write body", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.Linker.GetAllStatistics(r.Context())
	if err != nil {
		slog.Info("GetStatistics", "err service call", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(data)
	if err != nil {
		slog.Info("GetStatistics", "err write body", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetStatisticsById(w http.ResponseWriter, r *http.Request) {
	uuid, err := getIdFromRequest(r)
	if err != nil {
		slog.Info("GetStatisticsByShortLink", "err get id", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	data, err := h.service.Linker.GetStatisticsById(r.Context(), uuid)
	if err != nil {
		slog.Info("GetStatisticsByShortLink", "err service call", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(data)
	if err != nil {
		slog.Info("GetStatisticsByShortLink", "err write body", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func getIdFromRequest(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	if uuid == "" {
		return "", core.ErrUuidNotFound
	}

	return uuid, nil
}
