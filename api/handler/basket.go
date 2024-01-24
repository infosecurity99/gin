package handler

import (
	"fmt"
	"net/http"
	"test/api/models"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateBasket(c *gin.Context) {
	createBasket := models.CreateBasket{}

	if err := c.ShouldBindJSON(&createBasket); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pkey, err := h.storage.Basket().CreateBasket(createBasket)
	if err != nil {
		handleResponse(c, "error while creating user", http.StatusInternalServerError, err)
		return
	}

	res, err := h.storage.Basket().GetBasketByID(models.PrimaryKey{ID: pkey})
	if err != nil {
		handleResponse(c, "error while getting user by id", http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, " ", http.StatusCreated, res)
}

func (h Handler) GetBasketByID(c *gin.Context) {

	var err error

	uid := c.Param("id")
	fmt.Println("uid", uid)
	basket, err := h.storage.Basket().GetBasketByID(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting basket by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, basket)
}

func (h Handler) GetBasketList(c *gin.Context) {

}

func (h Handler) UpdateBasket(c *gin.Context) {

}

func (h Handler) DeleteBasket(c *gin.Context) {

}
