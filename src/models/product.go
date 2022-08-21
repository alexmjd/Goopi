package models

import (
	"database/sql"
	"errors"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=$1", p.Id).Scan(&p.Name, &p.Price)
}

func (p *Product) createProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}
