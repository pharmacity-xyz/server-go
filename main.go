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
	sr.Route.Route(config.BASICAPI, func(r chi.Router) {
		r.Post("/product", productC.Add)
		r.Get("/product", productC.GetAll)
		r.Get("/product/{productId}", productC.GetProductByProductId)
		r.Get("/product/category/{categoryId}", productC.GetProductByCategoryId)
		r.Get("/product/search/{searchWord}", productC.Search)
		r.Get("/product/featured", productC.FeaturedProducts)
		r.Put("/product", productC.Update)
		r.Delete("/product/{productId}", productC.Delete)
	})
}

func (sr ServiceRouter) CartItemRouter(cartItemService *models.CartItemService) {
	cartItemC := controllers.CartItems{
		CartItemService: cartItemService,
	}
	sr.Route.Post(config.BASICAPI+`/cart/add`, cartItemC.Add)
	sr.Route.Get(config.BASICAPI+`/cart`, cartItemC.GetAll)
	sr.Route.Get(config.BASICAPI+`/cart/count`, cartItemC.Count)
	sr.Route.Put(config.BASICAPI+`/cart/update_quantity`, cartItemC.UpdateQuantity)
	sr.Route.Delete(config.BASICAPI+`/cart/{productId}`, cartItemC.Delete)
}

func (sr ServiceRouter) PaymentRouter(cartItemService *models.CartItemService, userService *models.UserService) {
	paymentC := controllers.Payments{
		CartItemService: cartItemService,
		UserService:     userService,
	}
	sr.Route.Post(config.BASICAPI+`/payment/checkout`, paymentC.CreateCheckoutSession)
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
	cartItemService := models.CartItemService{
		DB: db,
	}

	serviceRouter.AuthRouter(&userService)
	serviceRouter.UserRouter(&userService)
	serviceRouter.CategoryRouter(&categoryService)
	serviceRouter.ProductRouter(&productService)
	serviceRouter.CartItemRouter(&cartItemService)
	serviceRouter.PaymentRouter(&cartItemService, &userService)

	fmt.Println("Starting the server on :8000...")
	http.ListenAndServe(":8000", r)
}
