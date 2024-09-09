package product

import (
	"database/sql"

	"github.com/dickyanth/eco-bite-v1/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetProducts() ([]types.Product, error){
	rows, err := s.db.Query("SELECT * FROM product")
	if err != nil {
		return nil,err
	}

	products := make([]types.Product, 0)
	for rows.Next(){
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error){
	product := new(types.Product)

	err := rows.Scan(
		&product.Id,
		&product.Image,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Quantity,
		&product.Created_at,
	)

	if err != nil {
		return nil, err
	}

	return product, nil	
}