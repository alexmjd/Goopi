package models

import (
	"database/sql"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) GetProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=?;", p.Id).Scan(&p.Name, &p.Price)
}

func (p *Product) GetAllProduct(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM products;").Scan(&p)
}

func (p *Product) CreateProduct(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO products(name, price) VALUES(?, ?);", p.Name, p.Price).Scan(&p.Id)

	if err == sql.ErrNoRows {
		err = db.QueryRow("SELECT LAST_INSERT_ID();").Scan(&p.Id)

		if err != nil {
			return err
		}
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}

func (p *Product) UpdateProduct(db *sql.DB) error {
	_, err := db.Exec("UPDATE products SET name=?, price=? WHERE id=?", p.Name, p.Price, p.Id)

	return err
}

func (p *Product) DeleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=?", p.Id)

	return err
}
