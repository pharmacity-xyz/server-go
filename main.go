package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
