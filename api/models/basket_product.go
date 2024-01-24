package models

type BasketProduct struct {
	ID         string `json:"id"`
	Basket_id  string `json:"basket_id"`
	Product_Id string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

type CreateBasketProduct struct {
	Basket_ID  string `json:"basket_id"`
	Product_id string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

type UpdateBasketProduct struct {
	ID         string `json:"id"`
	Basket_Id  string `json:"basket_id"`
	Product_Id string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

type ResponseBasketProduct struct {
	BasketProducts []BasketProduct `json:"basketsproducts"`
	Count          int             `json:"count"`
}

//finish
