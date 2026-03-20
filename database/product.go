package database

var products []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

func Store(p Product) Product {
	p.ID = len(products) + 1
	products = append(products, p)
	return p
}

func List() []Product {
	return products
}

func Get(productId int) *Product {
	for _, p := range products {
		if p.ID == productId {
			return &p
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range products {
		if p.ID == product.ID {
			products[idx] = product
		}
	}
}

func Delete(productId int) {
	var tempList []Product 

	for _, p := range products {
		if p.ID != productId {
			tempList = append(tempList, p)
		}
	}
	products = tempList
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
	products = append(products, product1, product2, product3, product4, product5)
}
