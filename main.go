package main

import (
	"gocommerce/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/", handlers.HomeHandler)
	r.Get("/product/{id}", handlers.ProductHandler)
	r.Get("/cart", handlers.CartHandler)
	r.Get("/checkout", handlers.CheckoutHandler)

	// HTMX endpoints
	// r.Post("/cart/add/{id}", handlers.AddToCartHandler)
	// r.Patch("/cart/item/{id}/{action}", handlers.UpdateCartItemHandler)
	// r.Delete("/cart/item/{id}", handlers.RemoveCartItemHandler)

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
