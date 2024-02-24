package product

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDBProductService_Register(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	productService := NewDBProductService(db)

	testProduct := Product{
		ManagerID:   1,
		Category:    "Electronics",
		Price:       "999.99",
		Name:        "Test Product",
		Description: "This is a test product",
		Size:        "Medium",
		ExpiredDate: time.Now().AddDate(0, 1, 0),
	}

	mock.ExpectExec("INSERT INTO product").WithArgs(
		testProduct.ManagerID,
		testProduct.Category,
		testProduct.Price,
		testProduct.Name,
		testProduct.Description,
		testProduct.Size,
		testProduct.ExpiredDate,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err = productService.Register(testProduct)
	assert.NoError(t, err, "Unexpected error during product registration")
}

func TestDBProductService_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	productService := NewDBProductService(db)

	testProductID := 1
	testUpdateFields := map[string]interface{}{
		"price":       "799.99",
		"description": "Updated description",
	}

	mock.ExpectExec("UPDATE product").WithArgs(
		testUpdateFields["price"],
		testUpdateFields["description"],
		testProductID,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err = productService.Update(testProductID, testUpdateFields)
	assert.NoError(t, err, "Unexpected error during product update")
	assert.NoError(t, mock.ExpectationsWereMet(), "Not all expectations were met")
}

func TestDBProductService_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	productService := NewDBProductService(db)

	testProductID := 1

	mock.ExpectQuery("SELECT id, manager_id, category, price, name, description, size, expired_date FROM product").WithArgs(
		testProductID,
	).WillReturnRows(sqlmock.NewRows([]string{"id", "manager_id", "category", "price", "name", "description", "size", "expired_date"}).
		AddRow(testProductID, 1, "Electronics", "999.99", "Test Product", "This is a test product", "Medium", time.Now().AddDate(0, 1, 0)))

	_, err = productService.Get(testProductID)
	assert.NoError(t, err, "Unexpected error during product retrieval")
	assert.NoError(t, mock.ExpectationsWereMet(), "Not all expectations were met")
}

func TestDBProductService_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	productService := NewDBProductService(db)

	testProductID := 1

	mock.ExpectExec("DELETE FROM product").WithArgs(
		testProductID,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err = productService.Delete(testProductID)
	assert.NoError(t, err, "Unexpected error during product deletion")
	assert.NoError(t, mock.ExpectationsWereMet(), "Not all expectations were met")
}
