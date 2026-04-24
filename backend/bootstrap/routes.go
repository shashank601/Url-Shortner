package bootstrap

import (
	"net/http"

	"github.com/shashank601/url-shortner/backend/internals/middleware"
)

func InitRouter(dep *Dependencies) *http.ServeMux {
	mux := http.NewServeMux()

	// URL routes
	mux.Handle("POST /shorten", middleware.JWTAuth(http.HandlerFunc(dep.UrlHandler.ShortenUrl)))
	mux.HandleFunc("GET /{code}", dep.UrlHandler.GetUrl)

	// Auth routes
	mux.HandleFunc("POST /signup", dep.AuthHandler.Register)
	mux.HandleFunc("POST /login", dep.AuthHandler.Login)

	// Analytics routes
	mux.Handle("GET /analytics/{code}", middleware.JWTAuth(http.HandlerFunc(dep.AnalyticsHandler.GetAnalytics)))
	
	// User routes
	mux.Handle("GET /urls", middleware.JWTAuth(http.HandlerFunc(dep.UrlHandler.ListUserURLs)))

	mux.Handle("GET /verify", middleware.JWTAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})))

	return mux
}
