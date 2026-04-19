package handler

import (
	"encoding/json"
	"net/http"
	"github.com/shashank601/url-shortner/backend/internals/service"
	"github.com/shashank601/url-shortner/backend/internals/dto"
)

type UrlHandler struct {
	Service *service.UrlService
}

func NewUrlHandler(s *service.UrlService) *UrlHandler {
	return &UrlHandler{Service: s}
}

func (h *UrlHandler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var req dto.UrlShortenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	response, err := h.Service.ShortenUrl(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func (h *UrlHandler) GetUrl(w http.ResponseWriter, r *http.Request) {
	var req dto.GetUrlRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	response, err := h.Service.GetUrl(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	
}
