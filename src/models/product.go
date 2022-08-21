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

func (p *Product) GetProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=?", p.Id).Scan(&p.Name, &p.Price)
}

func (p *Product) CreateProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}
