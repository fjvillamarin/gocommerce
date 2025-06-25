package handlers

import (
	"fmt"
	"gocommerce/models"
	"gocommerce/templates"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// In-memory cart for demo purposes
var cart = models.Cart{Items: []models.CartItem{}}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.Home(models.Products, cart.GetItemCount())
	component.Render(r.Context(), w)
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
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

	component := templates.ProductDetail(*product, cart.GetItemCount())
	component.Render(r.Context(), w)
}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.Cart(cart)
	component.Render(r.Context(), w)
}

func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Printf("Adding to cart: Product=%s, Quantity=%d, Color=%s, Size=%s\n", product.Name, quantity, color, size)

	cart.AddItem(*product, quantity, color, size)

	// Check if this is an HTMX request
	if r.Header.Get("HX-Request") == "true" {
		// Return just the updated cart badge for HTMX
		w.Header().Set("Content-Type", "text/html")
		if cart.GetItemCount() > 0 {
			badge := fmt.Sprintf(`<span class="absolute -top-2 -right-2 bg-blue-600 text-white text-xs font-bold rounded-full h-5 w-5 flex items-center justify-center" id="cart-count">%d</span>`, cart.GetItemCount())
			w.Write([]byte(badge))
		}
	} else {
		// Regular form submission - redirect
		http.Redirect(w, r, fmt.Sprintf("/product/%d", id), http.StatusSeeOther)
	}
}

func UpdateCartItemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	action := chi.URLParam(r, "action")

	// Find the item in cart
	for i, item := range cart.Items {
		if item.Product.ID == id {
			if action == "increase" {
				cart.Items[i].Quantity++
			} else if action == "decrease" && item.Quantity > 1 {
				cart.Items[i].Quantity--
			} else if action == "decrease" && item.Quantity == 1 {
				cart.RemoveItem(id, item.Color, item.Size)
			}
			break
		}
	}

	// Return both cart content and order summary
	w.Header().Set("Content-Type", "text/html")
	
	// Send both components
	templates.CartContent(cart).Render(r.Context(), w)
	w.Write([]byte(`<div id="order-summary-wrapper" hx-swap-oob="true">`))
	templates.OrderSummary(cart).Render(r.Context(), w)
	w.Write([]byte(`</div>`))
}

func RemoveCartItemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Find and remove the item
	for _, item := range cart.Items {
		if item.Product.ID == id {
			cart.RemoveItem(id, item.Color, item.Size)
			break
		}
	}

	// Return both cart content and order summary
	w.Header().Set("Content-Type", "text/html")
	
	// Send both components
	templates.CartContent(cart).Render(r.Context(), w)
	w.Write([]byte(`<div id="order-summary-wrapper" hx-swap-oob="true">`))
	templates.OrderSummary(cart).Render(r.Context(), w)
	w.Write([]byte(`</div>`))
}

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.Checkout(cart)
	component.Render(r.Context(), w)
}
