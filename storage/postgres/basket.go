package postgres

import (
	"database/sql"
	"fmt"
	"test/api/models"
	"test/storage"

	"github.com/google/uuid"
)

type basketRepo struct {
	db *sql.DB
}

func NewBasketRepo(db *sql.DB) storage.IBasket {
	return &basketRepo{
		db: db,
	}
}

func (b *basketRepo) CreateBasket(bas models.CreateBasket) (string, error) {

	id := uuid.New()

	if _, err := b.db.Exec(`insert into baskets(id, customer_id, total_sum)
	values($1, $2, $3)`,
		id,
		bas.CustomerID,
		bas.TotalSum,
	); err != nil {
		fmt.Println("error is while inserting data", err.Error())
		return "", err
	}

	return id.String(), nil
}

func (b *basketRepo) GetBasketByID(key models.PrimaryKey) (models.Basket, error) {
	basket := models.Basket{}

	if err := b.db.QueryRow(`
	select id, customer_id, total_sum from baskets where id = $1`,
		key.ID).Scan(
		&basket.ID,
		&basket.CustomerID,
		&basket.TotalSum); err != nil {
		fmt.Println("error is while selecting basket", err.Error())
		return models.Basket{}, err
	}

	return basket, nil
}

func (b *basketRepo) GetBasketList(request models.GetListRequest) (models.BasketResponse, error) {

	return models.BasketResponse{}, nil
}

func (b *basketRepo) UpdateBasket(basket models.UpdateBasket) (string, error) {

	return "", nil
}

func (b *basketRepo) DeleteBasket(key models.PrimaryKey) error {

	return nil
}
