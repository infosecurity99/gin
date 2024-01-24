package handler

import (
	"errors"
	"net/http"
	"strconv"
	"test/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//create
func (h Handler) CreateBasketProduct(c *gin.Context) {
	createbasketproduct := models.CreateBasketProduct{}

	if err := c.ShouldBindJSON(&createbasketproduct); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pkey, err := h.storage.BasketProduct().CreateBasketProduct(createbasketproduct)
	if err != nil {
		handleResponse(c, "error while creating user", http.StatusInternalServerError, err)
		return
	}

	res, err := h.storage.BasketProduct().GetBasketProductById(models.PrimaryKey{ID: pkey})
	if err != nil {
		handleResponse(c, "error while getting basketproducts by id", http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, " ", http.StatusCreated, res)
}

//getby id
func (h Handler) GetBasketProductById(c *gin.Context) {
	var err error

	uid := c.Param("id")

	basketproduct, err := h.storage.BasketProduct().GetBasketProductById(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting basketproduct by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, basketproduct)
}

//get list
func (h Handler) GetListBasketProduct(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	resp, err := h.storage.BasketProduct().GetListBasketProduct(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting users", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

//update busketproduct
func (h Handler) UpdateBasketProduct(c *gin.Context) {
	updateBasketproducts := models.UpdateBasketProduct{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateBasketproducts.ID = uid

	if err := c.ShouldBindJSON(&updateBasketproducts); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.BasketProduct().UpdateBasketProduct(updateBasketproducts)
	if err != nil {
		handleResponse(c, "error while updating updatebasket", http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.storage.BasketProduct().GetBasketProductById(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting basket by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, user)
}

//delete busketproduct
func (h Handler) DeleteBasketProduct(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.BasketProduct().DeleteBasketProduct(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting bakset by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}
