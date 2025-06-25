package models

import (
	"sync"
	"time"
)

// CartService manages shopping carts for multiple sessions
type CartService struct {
	mu    sync.RWMutex
	carts map[string]*Cart // sessionID -> Cart
}

// NewCartService creates a new cart service instance
func NewCartService() *CartService {
	return &CartService{
		carts: make(map[string]*Cart),
	}
}

// GetCart retrieves or creates a cart for a session
func (cs *CartService) GetCart(sessionID string) *Cart {
	cs.mu.RLock()
	cart, exists := cs.carts[sessionID]
	cs.mu.RUnlock()

	if !exists {
		cs.mu.Lock()
		// Double-check after acquiring write lock
		cart, exists = cs.carts[sessionID]
		if !exists {
			cart = &Cart{Items: []CartItem{}}
			cs.carts[sessionID] = cart
		}
		cs.mu.Unlock()
	}

	return cart
}

// AddItem adds an item to a specific session's cart
func (cs *CartService) AddItem(sessionID string, product Product, quantity int, color, size string) {
	cart := cs.GetCart(sessionID)
	cart.AddItem(product, quantity, color, size)
}

// UpdateQuantity updates the quantity of an item in a cart
func (cs *CartService) UpdateQuantity(sessionID string, productID int, action string) bool {
	cart := cs.GetCart(sessionID)

	// Find the item in cart
	for i, item := range cart.Items {
		if item.Product.ID == productID {
			if action == "increase" {
				cart.Items[i].Quantity++
			} else if action == "decrease" && item.Quantity > 1 {
				cart.Items[i].Quantity--
			} else if action == "decrease" && item.Quantity == 1 {
				cart.RemoveItem(productID, item.Color, item.Size)
			}
			return true
		}
	}
	return false
}

// RemoveItem removes an item from a specific session's cart
func (cs *CartService) RemoveItem(sessionID string, productID int) bool {
	cart := cs.GetCart(sessionID)

	// Find and remove the item
	for _, item := range cart.Items {
		if item.Product.ID == productID {
			cart.RemoveItem(productID, item.Color, item.Size)
			return true
		}
	}
	return false
}

// GetItemCount returns the total item count for a session's cart
func (cs *CartService) GetItemCount(sessionID string) int {
	cart := cs.GetCart(sessionID)

	// Delay of 1 second
	time.Sleep(10 * time.Second)

	return cart.GetItemCount()
}

// CleanupOldCarts removes carts that haven't been accessed in a while
// This could be called periodically to prevent memory growth
func (cs *CartService) CleanupOldCarts() {
	// For now, this is a placeholder. In a real app, you'd track last access time
	// and remove carts older than X hours/days
}
