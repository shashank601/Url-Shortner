package handler

import (
	"encoding/json"
	"net/http"
	"strings"
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
	shortCode := r.PathValue("code")
	if shortCode == "" {
		http.Error(w, "missing code", http.StatusBadRequest)
		return
	}

	response, err := h.Service.GetUrl(r.Context(), dto.GetUrlRequest{
		ShortCode: shortCode,
		UserAgent: r.UserAgent(),
		Referer:   r.Referer(),
		IP:        h.getIP(r),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, response.OriginalUrl, http.StatusFound)
}



// getIP extracts the real client IP from request headers
func (h *UrlHandler) getIP(r *http.Request) string {
	if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
		parts := strings.Split(fwd, ",")
		if len(parts) > 0 {
			return strings.TrimSpace(parts[0])
		}
	}

	ip := r.Header.Get("X-Real-IP")
	if ip != "" {
		return strings.TrimSpace(ip)
	}

	return r.RemoteAddr
}