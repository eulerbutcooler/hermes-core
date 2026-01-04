package api

import (
	"encoding/json"
	"net/http"

	"github.com/eulerbutcooler/hermes-core/internal/models"
	"github.com/eulerbutcooler/hermes-core/internal/store"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store *store.RelayStore
}

func NewHandler(s *store.RelayStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) CreateRelay(w http.ResponseWriter, r *http.Request) {
	var req models.CreateRelayRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.ActionType == "" {
		http.Error(w, "Name and ActionType are required", http.StatusBadRequest)
		return
	}
	id, err := h.store.CreateRelay(r.Context(), req)
	if err != nil {
		http.Error(w, "Failed to create relay: "+err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":   "success",
		"relay_id": id,
	})
}

func (h *Handler) GetRelayLogs(w http.ResponseWriter, r *http.Request) {
	relayID := chi.URLParam(r, "relayID")
	logs, err := h.store.GetLogs(r.Context(), relayID)
	if err != nil {
		http.Error(w, "Failed to fetch logs", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(logs)
}
