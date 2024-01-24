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
	var (
		baskets           = []models.Basket{}
		count             = 0
		countQuery, query string
		page              = request.Page
		offset            = (page - 1) * request.Limit
	)

	countQuery = `
		SELECT count(1) from   baskets`

	if err := b.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of basket", err.Error())
		return models.BasketResponse{}, err
	}

	query = `
	select id, customer_id, total_sum from baskets 
			    `

	query += ` LIMIT $1 OFFSET $2`

	rows, err := b.db.Query(query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.BasketResponse{}, err
	}

	for rows.Next() {
		basket := models.Basket{}

		if err = rows.Scan(
			&basket.ID,
			&basket.CustomerID,
			&basket.TotalSum,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.BasketResponse{}, err
		}

		baskets = append(baskets, basket)
	}

	return models.BasketResponse{
		Baskets: baskets,
		Count:   count,
	}, nil

}

func (b *basketRepo) UpdateBasket(request models.UpdateBasket) (string, error) {
	query := `
	update baskets
		set customer_id= $1, total_sum = $2
			where id = $3`

	if _, err := b.db.Exec(query, request.CustomerID, request.TotalSum, request.ID); err != nil {
		fmt.Println("error while updating baskets data", err.Error())
		return "", err
	}

	return request.ID, nil

}

func (b *basketRepo) DeleteBasket(key models.PrimaryKey) error {
	query := `
		delete from baskets
			where id = $1
`
	if _, err := b.db.Exec(query, key.ID); err != nil {
		fmt.Println("error while deleting basket by id", err.Error())
		return err
	}

	return nil
}
