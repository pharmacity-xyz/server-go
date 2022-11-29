package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/pharmacity-xyz/server-go/config"
	"github.com/pharmacity-xyz/server-go/controllers"
	"github.com/pharmacity-xyz/server-go/models"
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

	r.Get(config.BASICAPI+`/healthcheck`, controllers.HealthCheck)

	userService := models.UserService{
		DB: db,
	}
	authC := controllers.Auths{
		UserService: &userService,
	}
	r.Post(config.BASICAPI+`/auth/register`, authC.Register)
	r.Post(config.BASICAPI+`/auth/login`, authC.Login)
	r.Post(config.BASICAPI+`/auth/change_password`, authC.ChangePassword)

	userC := controllers.Users{
		UserService: &userService,
	}
	r.Get(config.BASICAPI+`/user`, userC.GetAll)
	r.Put(config.BASICAPI+`/user`, userC.Update)

	categoryService := models.CategoryService{
		DB: db,
	}
	categoryC := controllers.Categories{
		CategoryService: &categoryService,
	}
	r.Get(config.BASICAPI+`/category`, categoryC.GetAll)
	r.Post(config.BASICAPI+`/category`, categoryC.Add)
	r.Put(config.BASICAPI+`/category`, categoryC.Update)
	r.Delete(config.BASICAPI+`/category/{categoryId}`, categoryC.Delete)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
