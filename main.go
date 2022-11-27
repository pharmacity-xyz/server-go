package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pharmacity-xyz/server-go/controllers"
	"github.com/pharmacity-xyz/server-go/models"
)

const (
	BASICAPI = "/api/v1"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get(BASICAPI+`/healthcheck`, controllers.HealthCheck)

	userService := models.UserService{}
	userC := controllers.Users{
		UserService: &userService,
	}
	r.Post(BASICAPI+`/auth/register`, userC.Register)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
