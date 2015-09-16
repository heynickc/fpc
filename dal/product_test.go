package dal

import (
	"fmt"
	"testing"
	"time"

	"github.com/heynickc/fpc/libstring"
	_ "github.com/lib/pq"
)

func newProductForTest(t *testing.T) *Product {
	return NewProduct(newDbForTest(t))
}

func newProductNameForTest() string {
	return fmt.Sprintf("product-%v", libstring.RandString(32))
}

func TestProductCRUD(t *testing.T) {
	p := newProductForTest(t)

	description := "Roasted coffee on 9-17-15. Quatemalan, Ethiopian, Columbian"
	roastDate := time.Now()
	price := 19.99

	productRow, err := p.InsertNewProduct(nil, newProductNameForTest(), description, roastDate, price)

	if err != nil {
		t.Errorf("Adding new product should work.  Error: %v", err)
	}
	if productRow == nil {
		t.Fatal("Adding a new product should work")
	}
	if productRow.ID <= 0 {
		t.Fatal("Adding a new product should work")
	}

	// DELETE FROM products WHERE id=...
	_, err = p.DeleteById(nil, productRow.ID)

	if err != nil {
		t.Fatalf("Deleting user by id should not fail. Error: %v", err)
	}
}

func TestAllProducts(t *testing.T) {
	p := newProductForTest(t)

	productRows := make([]*ProductRow, 0)

	for i := 0; i < 10; i++ {
		description := "Roasted Day: 9-17-15. Quatemalan, Ethiopian, Columbian"
		roastDate := time.Now()
		price := 19.99

		productRow, err := p.InsertNewProduct(nil, newProductNameForTest(), description, roastDate, price)
		productRows = append(productRows, productRow)

		if err != nil {
			t.Error("Could not insert records for get all test")
		}
	}

	products, err := p.AllProducts(nil)

	if err != nil {
		t.Errorf("Error getting all products.  Error: %v", err)
	}
	if len(products) != 10 {
		t.Errorf("Expected 10 products, but got %v", len(products))
	}

	// Delete all product rows that were inserted for test
	for _, productRow := range productRows {
		_, err = p.DeleteById(nil, productRow.ID)
	}
}

func TestProductRowFormattedDate(t *testing.T) {
	now := time.Now()
	p := &ProductRow{RoastDate: now}
	expected := now.Format("Monday Jan 02, 2006")

	if p.FormattedRoastDate() != expected {
		t.Errorf("Bad formatting of roast date. Expected %v, but got %v", expected, p.FormattedRoastDate())
	}
}
