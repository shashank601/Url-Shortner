package bootstrap

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/shashank601/url-shortner/backend/internals/middleware"
)

func InitRouter(dep *Dependencies) *http.ServeMux {
	mux := http.NewServeMux()

	// URL routes
	mux.Handle("POST /api/shorten", middleware.JWTAuth(http.HandlerFunc(dep.UrlHandler.ShortenUrl)))
	mux.HandleFunc("GET /api/{code}", dep.UrlHandler.GetUrl)

	// Auth routes
	mux.HandleFunc("POST /api/signup", dep.AuthHandler.Register)
	mux.HandleFunc("POST /api/login", dep.AuthHandler.Login)

	// Analytics routes
	mux.Handle("GET /api/analytics/{code}", middleware.JWTAuth(http.HandlerFunc(dep.AnalyticsHandler.GetAnalytics)))
	
	// User routes
	mux.Handle("GET /api/urls", middleware.JWTAuth(http.HandlerFunc(dep.UrlHandler.ListUserURLs)))

	mux.Handle("GET /api/verify", middleware.JWTAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})))

	

	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exeDir := filepath.Dir(exePath)
	clientDistPath := filepath.Join(exeDir, "..", "client", "dist")


	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(filepath.Join(clientDistPath, "assets")))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		if r.URL.Path == "/api" || strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}
		
		http.ServeFile(w, r, filepath.Join(clientDistPath, "index.html"))
	})
	return mux
}
