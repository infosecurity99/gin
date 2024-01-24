package models

type BasketProduct struct {
	ID         string
	Basket_id  string
	Product_Id string
	Quantity   int
}

type CreateBasketProduct struct {
	Basket_ID  string
	Product_id string
	Quantity   int
}

type UpdateBasketProduct struct {
	ID         string
	Basket_Id  string
	Product_Id string
	Quantity   int
}

type ResponseBasketProduct struct {
	BasketProducts []BasketProduct `json:"basketsproducts"`
	Count          int             `json:"count"`
}
