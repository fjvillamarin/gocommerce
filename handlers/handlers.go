package handlers

import (
	"fmt"
	"gocommerce/models"
	"gocommerce/templates"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Handlers holds all the dependencies for HTTP handlers
type Handlers struct {
	cartService *models.CartService
}

// NewHandlers creates a new Handlers instance with dependencies
func NewHandlers(cartService *models.CartService) *Handlers {
	return &Handlers{
		cartService: cartService,
	}
}

// getSessionID extracts the session ID from cookies
func getSessionID(r *http.Request) string {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie.Value == "" {
		// This shouldn't happen if our JS is working, but provide a fallback
		return "default-session"
	}
	return cookie.Value
}

func (h *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := getSessionID(r)
	cartCount := h.cartService.GetItemCount(sessionID)
	component := templates.Home(models.Products, cartCount)
	component.Render(r.Context(), w)
}

func (h *Handlers) ProductHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := getSessionID(r)
	cartCount := h.cartService.GetItemCount(sessionID)

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := models.GetProductByID(id)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	component := templates.ProductDetail(*product, cartCount)
	component.Render(r.Context(), w)
}

func (h *Handlers) CartHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := getSessionID(r)
	cart := h.cartService.GetCart(sessionID)
	component := templates.Cart(*cart)
	component.Render(r.Context(), w)
}

func (h *Handlers) AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := getSessionID(r)

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := models.GetProductByID(id)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Parse form data
	r.ParseForm()

	// Get form data
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	if quantity <= 0 {
		quantity = 1
	}
	color := r.FormValue("color")
	size := r.FormValue("size")

	// Debug log
	fmt.Printf("Session %s - Adding to cart: Product=%s, Quantity=%d, Color=%s, Size=%s\n", sessionID, product.Name, quantity, color, size)

	h.cartService.AddItem(sessionID, *product, quantity, color, size)

	// Check if this is an HTMX request
	if r.Header.Get("HX-Request") == "true" {
		// Return just the updated cart badge for HTMX
		w.Header().Set("Content-Type", "text/html")
		cartCount := h.cartService.GetItemCount(sessionID)
		if cartCount > 0 {
			badge := fmt.Sprintf(`<span class="absolute -top-2 -right-2 bg-blue-600 text-white text-xs font-bold rounded-full h-5 w-5 flex items-center justify-center" id="cart-count">%d</span>`, cartCount)
			w.Write([]byte(badge))
		}
	} else {
		// Regular form submission - redirect
		http.Redirect(w, r, fmt.Sprintf("/product/%d", id), http.StatusSeeOther)
	}
}

func (h *Handlers) UpdateCartItemHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := getSessionID(r)

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	action := chi.URLParam(r, "action")
	h.cartService.UpdateQuantity(sessionID, id, action)

	// Get updated cart
	cart := h.cartService.GetCart(sessionID)

	// Return both cart content and order summary
	w.Header().Set("Content-Type", "text/html")

	// Send both components
	templates.CartContent(*cart).Render(r.Context(), w)
	w.Write([]byte(`<div id="order-summary-wrapper" hx-swap-oob="true">`))
	templates.OrderSummary(*cart).Render(r.Context(), w)
	w.Write([]byte(`</div>`))
}

func (h *Handlers) RemoveCartItemHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := getSessionID(r)

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	h.cartService.RemoveItem(sessionID, id)

	// Get updated cart
	cart := h.cartService.GetCart(sessionID)

	// Return both cart content and order summary
	w.Header().Set("Content-Type", "text/html")

	// Send both components
	templates.CartContent(*cart).Render(r.Context(), w)
	w.Write([]byte(`<div id="order-summary-wrapper" hx-swap-oob="true">`))
	templates.OrderSummary(*cart).Render(r.Context(), w)
	w.Write([]byte(`</div>`))
}

func (h *Handlers) CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := getSessionID(r)
	cart := h.cartService.GetCart(sessionID)
	component := templates.Checkout(*cart)
	component.Render(r.Context(), w)
}
