package api

import (
	"test/api/handler"
	"test/storage"

	"github.com/gin-gonic/gin"
)

func New(store storage.IStorage) *gin.Engine {
	h := handler.New(store)

	r := gin.New()

	//  for users  fisnish
	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUser)
	r.GET("/users", h.GetUserList)
	r.PUT("/user/:id", h.UpdateUser)
	r.DELETE("/user/:id", h.DeleteUser)
	r.PATCH("/user/:id", h.UpdateUserPassword)

	//  for baskets   fisnish
	r.POST("/basket", h.CreateBasket)
	r.GET("/basket/:id", h.GetBasketByID)
	r.GET("/baskets", h.GetBasketList)
	r.PUT("/basket/:id", h.UpdateBasket)
	r.DELETE("/basket/:id", h.DeleteBasket)

	//for products  finish
	r.POST("/product", h.CreateProduct)
	r.GET("/product/:id", h.GetByIdProduct)
	r.GET("/products", h.GetListProduct)
	r.PUT("/product/:id", h.UpdateProduct)
	r.DELETE("/product/:id", h.DelateProduct)

	//  for categories  finish
	r.POST("/category", h.CreateCategories)
	r.GET("/category/:id", h.GetCategoriesById)
	r.GET("/categorys", h.GetCategoriesGetList)
	r.PUT("/category/:id", h.UpdateCategories)
	r.DELETE("/category/:id", h.DelateCategories)

	//basketproducts
	r.POST("/basket_product", h.CreateBasketProduct)
	r.GET("/basket_product/:id", h.GetBasketProductById)
	r.GET("/basket_products", h.GetListBasketProduct)
	r.PUT("/basket_product/:id", h.UpdateBasketProduct)
	r.DELETE("/basket_product/:id", h.DeleteBasketProduct)

	return r
}
