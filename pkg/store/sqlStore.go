package store

import (
	"database/sql"
	"log"

	"github.com/AlejaMarin/go-api/internal/domain"
)

type sqlStore struct {
	DB *sql.DB
}

func NewSqlStore(database *sql.DB) StoreInterface {
	return &sqlStore{
		DB: database,
	}
}

func (s *sqlStore) GetAll() ([]domain.Product, error) {

	var p domain.Product
	var products []domain.Product
	query := "SELECT * FROM products;"
	rows, err := s.DB.Query(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&p.Id, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		} else {
			products = append(products, p)
		}
	}
	return products, nil
}

func (s *sqlStore) Read(id int) (domain.Product, error) {

	var product domain.Product

	query := "SELECT * FROM products WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *sqlStore) Create(product domain.Product) error {

	query := "INSERT INTO products(name, quantity, code_value, is_published, expiration, price) VALUES(?, ?, ?, ?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Update(product domain.Product) error {

	query := "UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?;"

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price, product.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Delete(id int) error {

	query := "DELETE FROM products WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Exists(codeValue string) bool {

	query := "SELECT id FROM products WHERE code_value = ?"
	row := s.DB.QueryRow(query, codeValue)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return false
	}

	if id > 0 {
		return true
	}
	return false
}
