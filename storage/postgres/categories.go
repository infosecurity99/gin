package postgres

import (
	"database/sql"
	"fmt"
	"test/api/models"
	"test/storage"

	"github.com/google/uuid"
)

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) storage.ICategory {
	return categoryRepo{
		db: db,
	}
}

//create category
func (c categoryRepo) CreateCategory(category models.CreateCategory) (string, error) {
	uid := uuid.New()

	if _, err := c.db.Exec(`insert into   categories  values ($1, $2)`,
		uid,
		category.Name,
	); err != nil {
		fmt.Println("error while inserting data  categories", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//get by idcategory
func (c categoryRepo) GetByIdCategory(pKey models.PrimaryKey) (models.Category, error) {
	category := models.Category{}

	query := `
		select id , name from  categories where id=$1 
`
	if err := c.db.QueryRow(query, pKey.ID).Scan(
		&category.ID,
		&category.Name,
	); err != nil {
		fmt.Println("error while scanning category", err.Error())
		return models.Category{}, err
	}

	return category, nil
}

//get  list

func (c categoryRepo) GetListCategory(request models.GetListRequest) (models.CategoriesResponse, error) {
	var (
		categories        = []models.Category{}
		count             = 0
		countQuery, query string
		page              = request.Page
		offset            = (page - 1) * request.Limit
	)

	countQuery = `select  count(1) from categories  `

	if err := c.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of categories", err.Error())
		return models.CategoriesResponse{}, err
	}

	query = `select  id, name  from categories `

	query += ` LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.CategoriesResponse{}, err
	}

	for rows.Next() {
		category := models.Category{}

		if err = rows.Scan(
			&category.ID,
			&category.Name,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.CategoriesResponse{}, err
		}

		categories = append(categories, category)
	}

	return models.CategoriesResponse{
		Categories: categories,
		Count:      count,
	}, nil
}

//update
func (c categoryRepo) UpdateCategory(request models.UpdateCategory) (string, error) {
	query := `
		update categories 
			set name = $1  where  id = $2`

	if _, err := c.db.Exec(query, request.Name, request.ID); err != nil {
		fmt.Println("error while updating categories data", err.Error())
		return "", err
	}

	return request.ID, nil
}

//delete

func (c categoryRepo) DeleteCategory(request models.PrimaryKey) error {
	query := `
		delete from categories
			where id = $1
`
	if _, err := c.db.Exec(query, request.ID); err != nil {
		fmt.Println("error while deleting categories by id", err.Error())
		return err
	}

	return nil
}
