package storage

import (
	"test/api/models"
)

type IStorage interface {
	Close()
	User() IUserStorage
	Basket() IBasket
	Category() ICategory
	Product() IProduct
	BasketProduct() IBasketProduct
}

//  for  user interface
type IUserStorage interface {
	Create(models.CreateUser) (string, error)
	GetByID(models.PrimaryKey) (models.User, error)
	GetList(models.GetListRequest) (models.UsersResponse, error)
	Update(models.UpdateUser) (string, error)
	Delete(models.PrimaryKey) error
	GetPassword(id string) (string, error)
	UpdatePassword(password models.UpdateUserPassword) error
}

//    basket for interface
type IBasket interface {
	CreateBasket(models.CreateBasket) (string, error)
	GetBasketByID(models.PrimaryKey) (models.Basket, error)
	GetBasketList(models.GetListRequest) (models.BasketResponse, error)
	UpdateBasket(models.UpdateBasket) (string, error)
	DeleteBasket(key models.PrimaryKey) error
}

//category fotr interfyce
type ICategory interface {
	CreateCategory(models.CreateCategory) (string, error)
	GetByIdCategory(models.PrimaryKey) (models.Category, error)
	GetListCategory(models.GetListRequest) (models.CategoriesResponse, error)
	UpdateCategory(models.UpdateCategory) (string, error)
	DeleteCategory(key models.PrimaryKey) error
}

//product for interface
type IProduct interface {
	CreateProduct(models.CreateProduct) (string, error)
	GetByIdProduct(models.PrimaryKey) (models.Product, error)
	GetListProduct(models.GetListRequest) (models.ProductResponce, error)
	UpdateProduct(models.UpdateProduct) (string, error)
	DelateProduct(models.PrimaryKey) error
}

//product   busket   for interface
type IBasketProduct interface {
	CreateBasketProduct(models.CreateBasketProduct) (string, error)
	GetBasketProductById(models.PrimaryKey) (models.BasketProduct, error)
	GetListBasketProduct(models.GetListRequest) (models.ResponseBasketProduct, error)
	UpdateBasketProduct(models.UpdateBasketProduct) (string, error)
	DeleteBasketProduct(models.PrimaryKey) error
}
