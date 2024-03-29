package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test/api/models"
	"test/storage"
)

type Handler struct {
	storage storage.IStorage
}

func New(store storage.IStorage) Handler {
	return Handler{
		storage: store,
	}
}

func handleResponse(c *gin.Context, msg string, statusCode int, data interface{}) {
	resp := models.Response{}

	switch code := statusCode; {
	case code < 400:
		resp.Description = "success"
	case code < 500:
		resp.Description = "bad request"
		fmt.Println("BAD REQUEST: "+msg, "reason: ", data)
	default:
		resp.Description = "internal server error"
		fmt.Println("INTERNAL SERVER ERROR: "+msg, "reason: ", data)
	}
	fmt.Println("data: ", data)

	resp.StatusCode = statusCode
	resp.Data = data
	fmt.Println("resp ", resp)

	c.JSON(resp.StatusCode, resp)
}
