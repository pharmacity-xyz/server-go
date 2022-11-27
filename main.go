package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/pharmacity-xyz/server-go/controllers"
	"github.com/pharmacity-xyz/server-go/models"
)

const (
	BASICAPI = "/api/v1"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := models.Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get(BASICAPI+`/healthcheck`, controllers.HealthCheck)

	userService := models.UserService{
		DB: db,
	}
	userC := controllers.Users{
		UserService: &userService,
	}
	r.Post(BASICAPI+`/auth/register`, userC.Register)
	r.Post(BASICAPI+`/auth/login`, userC.Login)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
