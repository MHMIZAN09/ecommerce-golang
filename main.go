package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

var products []Product

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	preflightHandler(w, r)
	if r.Method != "GET" {
		http.Error(w, "Please give me get method", 400)
		return
	}
	sendData(w, products, 200)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	preflightHandler(w, r)
	if r.Method != "POST" {
		http.Error(w, "Please give me post method", 400)
		return
	}
	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request body", 400)
		return
	}
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	sendData(w, newProduct, 201)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
func main() {
	mux := http.NewServeMux() // router

	mux.HandleFunc("/hello", helloHandler) // route

	mux.HandleFunc("/products", getProductsHandler)          // route
	mux.HandleFunc("/products/create", createProductHandler) // route
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

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
