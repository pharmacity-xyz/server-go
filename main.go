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

func (sr ServiceRouter) UserRouter(userService *models.UserService) {
	userC := controllers.Users{
		UserService: userService,
	}
	sr.Route.Get(config.BASICAPI+`/user`, userC.GetAll)
	sr.Route.Put(config.BASICAPI+`/user`, userC.Update)
}

func (sr ServiceRouter) CategoryRouter(categoryService *models.CategoryService) {
	categoryC := controllers.Categories{
		CategoryService: categoryService,
	}
	sr.Route.Get(config.BASICAPI+`/category`, categoryC.GetAll)
	sr.Route.Post(config.BASICAPI+`/category`, categoryC.Add)
	sr.Route.Put(config.BASICAPI+`/category`, categoryC.Update)
	sr.Route.Delete(config.BASICAPI+`/category/{categoryId}`, categoryC.Delete)
}

func (sr ServiceRouter) ProductRouter(productService *models.ProductService) {
	productC := controllers.Products{
		ProductService: productService,
	}
	sr.Route.Post(config.BASICAPI+"/product", productC.Add)
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
	categoryService := models.CategoryService{
		DB: db,
	}
	productService := models.ProductService{
		DB: db,
	}

	serviceRouter.AuthRouter(&userService)
	serviceRouter.UserRouter(&userService)
	serviceRouter.CategoryRouter(&categoryService)
	serviceRouter.ProductRouter(&productService)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
