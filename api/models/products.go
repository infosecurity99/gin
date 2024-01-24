package models

type Product struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Price          int    `json:"price"`
	Original_Price int    `json:"original_price"`
	Quantity       int    `json:"quantity"`
	Category_Id    string `json:"category_id"`
}

type CreateProduct struct {
	Name           string ` json:"name"`
	Price          int    `json:"price"`
	Original_Price int    `json:"original_price"`
	Quantity       int    `json:"quantity"`
	Category_Id    string `json:"category_id"`
}

type UpdateProduct struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Price          int    `json:"price"`
	Original_Price int    `json:"original_price"`
	Quantity       int    `json:"quantity"`
	Category_Id    int    `json:"category_id"`
}

type ProductResponce struct {
	Products []Product `json:"prodcuts"`
	Count    int       `json:"count"`
}
