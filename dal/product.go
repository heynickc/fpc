package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

func NewProduct(db *sqlx.DB) *Product {
	product := &Product{}
	product.db = db
	product.table = "products"
	product.hasID = true

	return product
}

type ProductRow struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	RoastDate   time.Time `db:"roast_date"`
	Description string    `db:"description"`
	Price       float32   `db:"price"`
}

func (pr ProductRow) FormattedRoastDate() string {
	return pr.RoastDate.Format("Monday Jan 02, 2006")
}

type Product struct {
	Base
}

func (p *Product) productRowFromResult(tx *sqlx.Tx, sqlResult sql.Result) (*ProductRow, error) {
	productId, err := sqlResult.LastInsertId()
	if err != nil {
		return nil, err
	}

	return p.GetById(tx, productId)
}

// GetById returns record by id.
func (p *Product) GetById(tx *sqlx.Tx, id int64) (*ProductRow, error) {
	product := &ProductRow{}
	query := fmt.Sprintf("SELECT * FROM %v WHERE id=$1", p.table)
	err := p.db.Get(product, query, id)

	return product, err
}

func (p *Product) InsertNewProduct(tx *sqlx.Tx, name, description string, roastDate time.Time, price float64) (*ProductRow, error) {
	if name == "" {
		return nil, errors.New("Name cannot be blank.")
	}
	if description == "" {
		return nil, errors.New("Description cannot be blank.")
	}
	if price <= 0 {
		return nil, errors.New("Price is invalid.")
	}

	data := make(map[string]interface{})
	data["name"] = name
	data["description"] = description
	data["roast_date"] = roastDate
	data["price"] = price

	sqlResult, err := p.InsertIntoTable(tx, data)
	if err != nil {
		return nil, err
	}

	return p.productRowFromResult(tx, sqlResult)
}

// AllProducts returns all product rows.
func (p *Product) AllProducts(tx *sqlx.Tx) ([]*ProductRow, error) {
	products := []*ProductRow{}
	query := fmt.Sprintf("SELECT * FROM %v", p.table)
	err := p.db.Select(&products, query)

	return products, err
}
