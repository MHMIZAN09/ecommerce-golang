package database

var Products []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

func init() {
	product1 := Product{
		ID:          1,
		Title:       "Mango",
		Description: "A sweet and juicy tropical fruit.",
		Price:       2.99,
		ImgUrl:      "https://example.com/mango.jpg",
	}
	product2 := Product{
		ID:          2,
		Title:       "Banana",
		Description: "A sweet and juicy tropical fruit.",
		Price:       1.99,
		ImgUrl:      "https://example.com/banana.jpg",
	}
	product3 := Product{
		ID:          3,
		Title:       "Apple",
		Description: "A sweet and juicy temperate fruit.",
		Price:       3.99,
		ImgUrl:      "https://example.com/apple.jpg",
	}
	product4 := Product{
		ID:          4,
		Title:       "Pineapple",
		Description: "A sweet and juicy tropical fruit.",
		Price:       4.99,
		ImgUrl:      "https://example.com/pineapple.jpg",
	}
	product5 := Product{
		ID:          5,
		Title:       "Grapes",
		Description: "A sweet and juicy temperate fruit.",
		Price:       2.49,
		ImgUrl:      "https://example.com/grapes.jpg",
	}
	Products = append(Products, product1, product2, product3, product4, product5)
}
