package api

import (
	"test/api/handler"
	"test/storage"

	"github.com/gin-gonic/gin"
)

func New(store storage.IStorage) *gin.Engine {
	h := handler.New(store)

	r := gin.New()

	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUser)
	r.GET("/users", h.GetUserList)
	r.PUT("/user/:id", h.UpdateUser)
	r.DELETE("/user/:id", h.DeleteUser)
	r.PATCH("/user/:id", h.UpdateUserPassword)

	r.POST("/basket", h.CreateBasket)
	r.GET("/basket/:id", h.GetBasketByID)
	r.GET("/baskets", h.GetBasketList)
	r.PUT("/basket/:id", h.UpdateBasket)
	r.DELETE("/basket/:id", h.DeleteBasket)

	return r
}
