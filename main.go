package main

import (
	"gocommerce/handlers"
	"gocommerce/models"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize services
	cartService := models.NewCartService()
	
	// Initialize handlers with dependencies
	h := handlers.NewHandlers(cartService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/", h.HomeHandler)
	r.Get("/product/{id}", h.ProductHandler)
	r.Get("/cart", h.CartHandler)
	r.Get("/checkout", h.CheckoutHandler)

	// HTMX endpoints
	r.Post("/cart/add/{id}", h.AddToCartHandler)
	r.Patch("/cart/item/{id}/{action}", h.UpdateCartItemHandler)
	r.Delete("/cart/item/{id}", h.RemoveCartItemHandler)
	
	// API endpoints
	r.Get("/api/cart-count", h.CartCountHandler)

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
