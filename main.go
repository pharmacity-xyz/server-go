package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
	sr.Route.Route(config.BASICAPI+`/user`, func(r chi.Router) {
		r.Get(`/`, userC.GetAll)
		r.Put(`/`, userC.Update)
	})
}

func (sr ServiceRouter) CategoryRouter(categoryService *models.CategoryService) {
	categoryC := controllers.Categories{
		CategoryService: categoryService,
	}
	sr.Route.Route(config.BASICAPI+`/category`, func(r chi.Router) {
		r.Get(`/`, categoryC.GetAll)
		r.Post(`/`, categoryC.Add)
		r.Put(`/`, categoryC.Update)
		r.Delete(`/{categoryId}`, categoryC.Delete)
	})
}

func (sr ServiceRouter) ProductRouter(productService *models.ProductService) {
	productC := controllers.Products{
		ProductService: productService,
	}
	sr.Route.Route(config.BASICAPI+`/product`, func(r chi.Router) {
		r.Post("/", productC.Add)
		r.Get("/", productC.GetAll)
		r.Get("/{productId}", productC.GetProductByProductId)
		r.Get("/category/{categoryId}", productC.GetProductByCategoryId)
		r.Get("/search/{searchWord}", productC.Search)
		r.Get("/featured", productC.FeaturedProducts)
		r.Put("/", productC.Update)
		r.Delete("/{productId}", productC.Delete)
	})
}

func (sr ServiceRouter) CartItemRouter(cartItemService *models.CartItemService) {
	cartItemC := controllers.CartItems{
		CartItemService: cartItemService,
	}
	sr.Route.Route(config.BASICAPI+`/cart`, func(r chi.Router) {
		r.Post(`/add`, cartItemC.Add)
		r.Get(`/`, cartItemC.GetAll)
		r.Get(`/count`, cartItemC.Count)
		r.Put(`/update_quantity`, cartItemC.UpdateQuantity)
		r.Delete(`/{productId}`, cartItemC.Delete)
	})
}

func (sr ServiceRouter) OrderRouter(orderService *models.OrderService, categoryService *models.CategoryService) {
	orderC := controllers.Orders{
		CategoryService: categoryService,
		OrderService:    orderService,
	}
	sr.Route.Route(config.BASICAPI+`/order`, func(r chi.Router) {
		r.Get(`/`, orderC.GetOrders)
		r.Get(`/{orderId}`, orderC.GetOrderDetails)
		r.Get(`/admin`, orderC.GetOrdersForAdmin)
		r.Get(`/charts`, orderC.GetOrdersPerMonth)
		r.Get(`/piecharts`, orderC.GetOrdersForPieChart)
	})
}

func (sr ServiceRouter) PaymentRouter(
	cartItemService *models.CartItemService,
	userService *models.UserService,
	paymentService *models.PaymentService,
	orderService *models.OrderService,
) {
	paymentC := controllers.Payments{
		CartItemService: cartItemService,
		UserService:     userService,
		PaymentService:  paymentService,
		OrderService:    orderService,
	}
	sr.Route.Route(config.BASICAPI+`/payment`, func(r chi.Router) {
		r.Post(`/checkout`, paymentC.CreateCheckoutSession)
		r.Post(`/`, paymentC.FulfilOrder)
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
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

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
	paymentService := models.PaymentService{
		DB: db,
	}
	orderService := models.OrderService{
		DB: db,
	}

	serviceRouter.AuthRouter(&userService)
	serviceRouter.UserRouter(&userService)
	serviceRouter.CategoryRouter(&categoryService)
	serviceRouter.ProductRouter(&productService)
	serviceRouter.CartItemRouter(&cartItemService)
	serviceRouter.PaymentRouter(&cartItemService, &userService, &paymentService, &orderService)
	serviceRouter.OrderRouter(&orderService, &categoryService)

	mode := os.Getenv("MODE")
	fmt.Printf("Starting the server on :8000... %v", mode)

	if mode == "PRODUCTION" {
		http.ListenAndServe(":0000", r)
	} else {
		http.ListenAndServe(":8000", r)
	}
}
