package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	
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
	q := r.URL.Query()
	customerID, err := strconv.Atoi(q.Get("customer_id"))
	if err != nil {
		http.Error(w, "invalid customer_id", http.StatusBadRequest)
		return
	}
	urlID, err := strconv.Atoi(q.Get("url_id"))
	if err != nil {
		http.Error(w, "invalid url_id", http.StatusBadRequest)
		return
	}

	req := dto.AnalyticsRequest{CustomerID: customerID, UrlID: urlID}

	result, err := h.Service.GetAnalytics(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
