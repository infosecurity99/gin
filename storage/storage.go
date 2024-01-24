package storage

import (
	"test/api/models"
)

type IStorage interface {
	Close()
	User() IUserStorage
	Basket() IBasket
	//Category() ICategoryStorage
}

type IUserStorage interface {
	Create(models.CreateUser) (string, error)
	GetByID(models.PrimaryKey) (models.User, error)
	GetList(models.GetListRequest) (models.UsersResponse, error)
	Update(models.UpdateUser) (string, error)
	Delete(models.PrimaryKey) error
	GetPassword(id string) (string, error)
	UpdatePassword(password models.UpdateUserPassword) error
}

//type ICategoryStorage interface {}

type IBasket interface {
	CreateBasket(models.CreateBasket) (string, error)
	GetBasketByID(models.PrimaryKey) (models.Basket, error)
	GetBasketList(models.GetListRequest) (models.BasketResponse, error)
	UpdateBasket(models.UpdateBasket) (string, error)
	DeleteBasket(key models.PrimaryKey) error
}
