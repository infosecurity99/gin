package postgres

import (
	"database/sql"
	"fmt"
	"test/api/models"
	"test/storage"

	"github.com/google/uuid"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) storage.IProduct {
	return &productRepo{
		db: db,
	}
}

//create
func (p *productRepo) CreateProduct(createProduct models.CreateProduct) (string, error) {
	uid := uuid.New()

	if _, err := p.db.Exec(`insert into 
			products values ($1, $2, $3, $4, $5, $6)
			`,
		uid,
		createProduct.Name,
		createProduct.Price,
		createProduct.Original_Price,
		createProduct.Quantity,
		createProduct.Category_Id,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//getbyid

func (p *productRepo) GetByIdProduct(pKey models.PrimaryKey) (models.Product, error) {
	products := models.Product{}

	query := `
		select id, name, price, orginal_price,quantity, category_id from products where id = $1
`
	if err := p.db.QueryRow(query, pKey.ID).Scan(
		&products.ID,
		&products.Name,
		&products.Price,
		&products.Original_Price,
		&products.Category_Id,
	); err != nil {
		fmt.Println("error while scanning user", err.Error())
		return models.Product{}, err
	}

	return products, nil
}

//get list product

func (p *productRepo) GetListProduct(request models.GetListRequest) (models.ProductResponce, error) {
	var (
		products          = []models.Product{}
		count             = 0
		countQuery, query string
		page              = request.Page
		offset            = (page - 1) * request.Limit
	)

	countQuery = `
		SELECT count(1) from products where  `

	if err := p.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of products", err.Error())
		return models.ProductResponce{}, err
	}

	query = `
		SELECT id, name , original_price, quantity,category_id
			FROM products
			  
			    `

	query += ` LIMIT $1 OFFSET $2`

	rows, err := p.db.Query(query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.ProductResponce{}, err
	}

	for rows.Next() {
		product := models.Product{}

		if err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Original_Price,
			&product.Category_Id,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.ProductResponce{}, err
		}

		products = append(products, product)
	}

	return models.ProductResponce{
		Products: products,
		Count:    count,
	}, nil
}

//update  product

func (p *productRepo) UpdateProduct(request models.UpdateProduct) (string, error) {
	query := `
	update products 
		set name = $1, price = $2, original_price = $3, quantity=$4, category_id=$5
			where id = $6`

	if _, err := p.db.Exec(query, request.Name, request.Price, request.Original_Price, request.Quantity, request.Category_Id, request.ID); err != nil {
		fmt.Println("error while updating user data", err.Error())
		return "", err
	}

	return request.ID, nil
}

//delete product

func (p *productRepo) DelateProduct(request models.PrimaryKey) error {
	query := `
		delete from products
			where id = $1
`
	if _, err := p.db.Exec(query, request.ID); err != nil {
		fmt.Println("error while deleting products by id", err.Error())
		return err
	}

	return nil
}
