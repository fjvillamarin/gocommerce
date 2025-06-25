package models

var Products = []Product{
	{
		ID:          1,
		Name:        "Minimal Watch",
		Description: "Elegant timepiece with clean design",
		Price:       299,
		Image:       "https://images.pexels.com/photos/267394/pexels-photo-267394.jpeg?auto=compress&cs=tinysrgb&w=400",
		Colors:      []string{"Black", "Blue", "Rose Gold"},
		Sizes:       []string{"38mm", "42mm", "44mm"},
	},
	{
		ID:          2,
		Name:        "Leather Bag",
		Description: "Premium leather craftsmanship",
		Price:       189,
		Image:       "https://images.pexels.com/photos/335257/pexels-photo-335257.jpeg?auto=compress&cs=tinysrgb&w=400",
		Colors:      []string{"Brown", "Black", "Tan"},
		Sizes:       []string{"Small", "Medium", "Large"},
	},
	{
		ID:          3,
		Name:        "Wireless Headphones",
		Description: "Crystal clear audio experience",
		Price:       149,
		Image:       "https://images.pexels.com/photos/1029896/pexels-photo-1029896.jpeg?auto=compress&cs=tinysrgb&w=400",
		Colors:      []string{"Black", "White", "Blue"},
		Sizes:       []string{"One Size"},
	},
	{
		ID:          4,
		Name:        "Ceramic Mug",
		Description: "Handcrafted ceramic design",
		Price:       39,
		Image:       "https://images.pexels.com/photos/341523/pexels-photo-341523.jpeg?auto=compress&cs=tinysrgb&w=400",
		Colors:      []string{"White", "Blue", "Green"},
		Sizes:       []string{"Regular", "Large"},
	},
	{
		ID:          5,
		Name:        "Modern Plant Pot",
		Description: "Stylish home decoration",
		Price:       59,
		Image:       "https://images.pexels.com/photos/1649771/pexels-photo-1649771.jpeg?auto=compress&cs=tinysrgb&w=400",
		Colors:      []string{"White", "Gray", "Terracotta"},
		Sizes:       []string{"Small", "Medium", "Large"},
	},
	{
		ID:          6,
		Name:        "Premium Notebook",
		Description: "High-quality paper and binding",
		Price:       29,
		Image:       "https://images.pexels.com/photos/1152077/pexels-photo-1152077.jpeg?auto=compress&cs=tinysrgb&w=400",
		Colors:      []string{"Black", "Brown", "Navy"},
		Sizes:       []string{"A5", "A4"},
	},
	{
		ID:          7,
		Name:        "Designer Glasses",
		Description: "Modern frame design",
		Price:       199,
		Image:       "https://images.pexels.com/photos/1407354/pexels-photo-1407354.jpeg?auto=compress&cs=tinysrgb&w=400",
		Colors:      []string{"Black", "Tortoise", "Clear"},
		Sizes:       []string{"Small", "Medium", "Large"},
	},
	{
		ID:          8,
		Name:        "Scented Candle",
		Description: "Natural soy wax blend",
		Price:       45,
		Image:       "https://images.pexels.com/photos/1464625/pexels-photo-1464625.jpeg?auto=compress&cs=tinysrgb&w=400",
		Colors:      []string{"Vanilla", "Lavender", "Citrus"},
		Sizes:       []string{"Small", "Large"},
	},
}

func GetProductByID(id int) *Product {
	for _, p := range Products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
