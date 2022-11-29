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

type ServiceRouter struct {
	Route *chi.Mux
}

func (sr ServiceRouter) HealchCheckRouter() {
	sr.Route.Get(config.BASICAPI+`/healthcheck`, controllers.HealthCheck)
}

func (sr ServiceRouter) AuthRouter(userService *models.UserService) {
	authC := controllers.Auths{
		UserService: userService,
	}
	sr.Route.Route(config.BASICAPI+`/auth`, func(r chi.Router) {
		r.Post(`/register`, authC.Register)
		r.Post(`/login`, authC.Login)
		r.Post(`/change_password`, authC.ChangePassword)
	})
}

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
	serviceRouter := ServiceRouter{Route: r}
	r.Use(middleware.Logger)

	serviceRouter.HealchCheckRouter()

	userService := models.UserService{
		DB: db,
	}

	serviceRouter.AuthRouter(&userService)

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
