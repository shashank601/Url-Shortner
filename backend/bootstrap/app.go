package bootstrap

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func StartApp() {
	_ = godotenv.Load()
	setDefaultEnv("PORT", "8080")

	port := os.Getenv("PORT")

	db := InitDB()
	rdb := InitRedis()

	deps := InitDependencies(db, rdb)
	router := InitRouter(deps)

	addr := ":" + port
	log.Println("server started on", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("server failed:", err)
	}
}

func setDefaultEnv(key, value string) {
	if os.Getenv(key) == "" {
		_ = os.Setenv(key, value)
	}
}
