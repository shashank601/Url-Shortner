package main

import (
	"fmt"
	"net/http"

	"github.com/shashank601/url-shortner/backend/internals/handler"
	"github.com/shashank601/url-shortner/backend/internals/repo"
	"github.com/shashank601/url-shortner/backend/internals/service"
)

func main() {
	userRepo := &repo.UserRepo{}

	userService := &service.UserService{Repo: userRepo}

	userHandler := &handler.UserHandler{Service: userService}

	http.HandleFunc("/users", userHandler.GetUsers)

	fmt.Println("server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server failed:", err)
	}
}
