package buyer

import (
	"database/sql"
	"fmt"

	"github.com/dickyanth/eco-bite-v1/types"
)

type Store struct{
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{db}
}

func (s *Store) GetBuyerByEmail(email string)(*types.Buyer, error){
	rows, err := s.db.Query("SELECT * FROM buyer WHERE email = ?", email)
	if err != nil{
		return nil, err
	}

	u := new(types.Buyer)
	for rows.Next(){
		u, err = scanRowIntoBuyer(rows)
		if err != nil{
			return nil, err
		}
	}
	if u.BuyerId == 0 {
		return nil, fmt.Errorf("buyer not found")
	}
	return u, nil
}

func scanRowIntoBuyer(rows *sql.Rows) (*types.Buyer, error){
	buyer := new(types.Buyer)

	err := rows.Scan(
		&buyer.BuyerId,
		&buyer.Name,
		&buyer.Email,
		&buyer.PasswordHash,
		&buyer.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return buyer,nil
}

func (s *Store) GetBuyerByID(id int) (*types.Buyer, error) {
	rows, err := s.db.Query("SELECT * FROM buyer WHERE id = ?", id)
	if err != nil{
		return nil, err
	}

	u := new(types.Buyer)
	for rows.Next(){
		u, err = scanRowIntoBuyer(rows)
		if err != nil{
			return nil, err
		}
	}
	if u.BuyerId == 0 {
		return nil, fmt.Errorf("buyer not found")
	}
	return u, nil
}

func (s *Store) CreateBuyer(buyer types.Buyer) error{
	_, err := s.db.Exec("INSERT INTO buyer (name, email, phone_number, address, password_hash) values (?, ?, ?, ?, ?)", buyer.Name, buyer.Email, buyer. PhoneNumber, buyer.Address, buyer.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}