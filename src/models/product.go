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

func GetAllProduct(db *sql.DB) ([]Product, error) {
	products := []Product{}

	rows, err := db.Query("SELECT id, name, price FROM products;")

	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (p *Product) CreateProduct(db *sql.DB) error {
	result, err := db.Exec("INSERT INTO products(name, price) VALUES(?, ?);", p.Name, p.Price)
	if err != nil {
		return err
	}

	productID, err := result.LastInsertId()
	p.Id = int(productID)

	if err != nil {
		return err
	}

	return nil
}

func (p *Product) UpdateProduct(db *sql.DB) error {
	_, err := db.Exec("UPDATE products SET name=?, price=? WHERE id=?;", p.Name, p.Price, p.Id)

	return err
}

func (p *Product) DeleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=?;", p.Id)

	return err
}
