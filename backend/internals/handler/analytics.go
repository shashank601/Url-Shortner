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
	shortCode := r.PathValue("code")
	if shortCode == "" {
		http.Error(w, "missing code", http.StatusBadRequest)
		return
	}

	customerID, ok := r.Context().Value("customer_id").(int)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	req := dto.AnalyticsRequest{
		CustomerID: customerID,
		ShortCode:  shortCode,
	}

	result, err := h.Service.GetAnalytics(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
