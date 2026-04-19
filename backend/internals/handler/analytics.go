package handler

import (
	"encoding/json"
	"net/http"
	
	"github.com/shashank601/url-shortner/backend/internals/dto"
	"github.com/shashank601/url-shortner/backend/internals/service"
)

type AnalyticsHandler struct {
	Service *service.AnalyticsService	
}

func NewAnalyticsHandler(service *service.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{
		Service: service,
	}
}

func (h *AnalyticsHandler) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	var req dto.AnalyticsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	result, err := h.Service.GetAnalytics(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
