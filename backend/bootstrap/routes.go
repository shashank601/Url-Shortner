package bootstrap

import "net/http"

func InitRouter(dep *Dependencies) *http.ServeMux {
	mux := http.NewServeMux()

	// URL routes
	mux.HandleFunc("POST /shorten", dep.UrlHandler.ShortenUrl)
	mux.HandleFunc("GET /{code}", dep.UrlHandler.GetUrl)

	// Auth routes
	mux.HandleFunc("POST /signup", dep.AuthHandler.Register)
	mux.HandleFunc("POST /login", dep.AuthHandler.Login)

	// Analytics routes
	mux.HandleFunc("GET /analytics/{code}", dep.AnalyticsHandler.GetAnalytics)

	return mux
}