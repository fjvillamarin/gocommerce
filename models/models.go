package models

type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Colors      []string `json:"colors"`
	Sizes       []string `json:"sizes"`
}

type CartItem struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
	Color    string  `json:"color"`
	Size     string  `json:"size"`
}

type Cart struct {
	Items []CartItem `json:"items"`
}

func (c *Cart) GetTotal() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.Product.Price * float64(item.Quantity)
	}
	return total
}

func (c *Cart) GetItemCount() int {
	count := 0
	for _, item := range c.Items {
		count += item.Quantity
	}
	return count
}

func (c *Cart) AddItem(product Product, quantity int, color, size string) {
	for i, item := range c.Items {
		if item.Product.ID == product.ID && item.Color == color && item.Size == size {
			c.Items[i].Quantity += quantity
			return
		}
	}
	c.Items = append(c.Items, CartItem{
		Product:  product,
		Quantity: quantity,
		Color:    color,
		Size:     size,
	})
}

func (c *Cart) RemoveItem(productID int, color, size string) {
	for i, item := range c.Items {
		if item.Product.ID == productID && item.Color == color && item.Size == size {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			return
		}
	}
}

func (c *Cart) UpdateQuantity(productID int, color, size string, quantity int) {
	for i, item := range c.Items {
		if item.Product.ID == productID && item.Color == color && item.Size == size {
			if quantity <= 0 {
				c.RemoveItem(productID, color, size)
			} else {
				c.Items[i].Quantity = quantity
			}
			return
		}
	}
}
