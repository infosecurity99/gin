package postgres

import (
	"database/sql"
	"fmt"
	"test/api/models"
	"test/storage"

	"github.com/google/uuid"
)

type basketproductRepo struct {
	db *sql.DB
}

func NewBasketProduct(db *sql.DB) storage.IBasketProduct {
	return &basketproductRepo{
		db: db,
	}
}

//create for interface   basketproduct
func (b *basketproductRepo) CreateBasketProduct(createbasketproduct models.CreateBasketProduct) (string, error) {
	uid := uuid.New()

	if _, err := b.db.Exec(`insert into 
			basket_products values ($1, $2, $3, $4)
			`,
		uid,
		createbasketproduct.Basket_ID,
		createbasketproduct.Product_id,
		createbasketproduct.Quantity,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//get by id  for interface basketproduct
func (b *basketproductRepo) GetBasketProductById(pKey models.PrimaryKey) (models.BasketProduct, error) {
	basketproduct := models.BasketProduct{}

	query := `
		select id, basket_id, product_id, quantity from basket_products where id = $1 
`
	if err := b.db.QueryRow(query, pKey.ID).Scan(
		&basketproduct.ID,
		&basketproduct.Basket_id,
		&basketproduct.Product_Id,
		&basketproduct.Quantity,
	); err != nil {
		fmt.Println("error while scanning user", err.Error())
		return models.BasketProduct{}, err
	}

	return basketproduct, nil
}

//get list  for  interface basketproduct
func (b *basketproductRepo) GetListBasketProduct(request models.GetListRequest) (models.ResponseBasketProduct, error) {
	var (
		basketproducts    = []models.BasketProduct{}
		count             = 0
		countQuery, query string
		page              = request.Page
		offset            = (page - 1) * request.Limit
	)

	countQuery = `
		SELECT count(1) from basket_products  `

	if err := b.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of users", err.Error())
		return models.ResponseBasketProduct{}, err
	}

	query = `
		SELECT id, basket_id, product_id, quantity
			FROM ba_products
			    `

	query += ` LIMIT $1 OFFSET $2`

	rows, err := b.db.Query(query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.ResponseBasketProduct{}, err
	}

	for rows.Next() {
		basketproduct := models.BasketProduct{}

		if err = rows.Scan(
			&basketproduct.ID,
			&basketproduct.Basket_id,
			&basketproduct.Product_Id,
			&basketproduct.Quantity,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.ResponseBasketProduct{}, err
		}

		basketproducts = append(basketproducts, basketproduct)
	}

	return models.ResponseBasketProduct{
		BasketProducts: basketproducts,
		Count:          count,
	}, nil
}

// update  for interface basketproduct
func (b *basketproductRepo) UpdateBasketProduct(request models.UpdateBasketProduct) (string, error) {
	query := `
		update basket_products 
			set basket_id = $1, product_id = $2, quantity = $3
				where  id = $4`

	if _, err := b.db.Exec(query, request.Basket_Id, request.Product_Id, request.Quantity, request.ID); err != nil {
		fmt.Println("error while updating user data", err.Error())
		return "", err
	}

	return request.ID, nil
}

//delete for interface basketproduct
func (b *basketproductRepo) DeleteBasketProduct(request models.PrimaryKey) error {
	query := `
		delete from basket_products
			where id = $1
`
	if _, err := b.db.Exec(query, request.ID); err != nil {
		fmt.Println("error while deleting user by id", err.Error())
		return err
	}

	return nil
}
