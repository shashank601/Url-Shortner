package bootstrap

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func StartApp() {
	_ = godotenv.Load()
	setDefaultEnv("PORT", "8080")

	port := os.Getenv("PORT")

	db := InitDB()
	rdb := InitRedis()

	deps := InitDependencies(db, rdb)
	router := InitRouter(deps)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // hardcoded for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(router)

	addr := ":" + port
	log.Println("server started on", addr)

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal("server failed:", err)
	}
}

func setDefaultEnv(key, value string) {
	if os.Getenv(key) == "" {
		_ = os.Setenv(key, value)
	}
}
