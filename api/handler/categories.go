package handler

import (
	"errors"
	"net/http"
	"strconv"
	"test/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// create
func (h Handler) CreateCategories(c *gin.Context) {
	createCategories := models.CreateCategory{}

	if err := c.ShouldBindJSON(&createCategories); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Category().CreateCategory(createCategories)
	if err != nil {
		handleResponse(c, "error while creating user", http.StatusInternalServerError, err)
		return
	}

	category, err := h.storage.Category().GetByIdCategory(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting category by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, category)
}

//get by id
func (h Handler) GetCategoriesById(c *gin.Context) {
	var err error

	uid := c.Param("id")

	user, err := h.storage.Category().GetByIdCategory(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting category by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, user)
}

// getlistcategory
func (h Handler) GetCategoriesGetList(c *gin.Context) {
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

	resp, err := h.storage.Category().GetListCategory(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting category", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

//updatecategory
func (h Handler) UpdateCategories(c *gin.Context) {
	updateCategory := models.UpdateCategory{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateCategory.ID = uid

	if err := c.ShouldBindJSON(&updateCategory); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Category().UpdateCategory(updateCategory)
	if err != nil {
		handleResponse(c, "error while updating category", http.StatusInternalServerError, err.Error())
		return
	}

	category, err := h.storage.Category().GetByIdCategory(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting user by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, category)
}

//delete category
func (h Handler) DelateCategories(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Category().DeleteCategory(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting delete by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}
