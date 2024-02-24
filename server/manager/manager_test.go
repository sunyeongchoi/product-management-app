package manager

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDBManagerService_SignUp(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	managerService := NewDBManagerService(db)

	testManager := Manager{
		Phone:    "1234567890",
		Password: "testpassword",
	}

	mock.ExpectExec("INSERT INTO manager").WithArgs(
		testManager.Phone,
		testManager.Password,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err = managerService.SignUp(testManager)
	assert.NoError(t, err, "Unexpected error during manager sign up")
	assert.NoError(t, mock.ExpectationsWereMet(), "Not all expectations were met")
}

func TestDBManagerService_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	managerService := NewDBManagerService(db)

	testPhone := "01012345678"

	mock.ExpectQuery("SELECT id, phone, password FROM manager").WithArgs(
		testPhone,
	).WillReturnRows(sqlmock.NewRows([]string{"id", "phone", "password"}).
		AddRow(1, testPhone, "testpassword"))

	_, err = managerService.Get(testPhone)
	assert.NoError(t, err, "Unexpected error during manager retrieval")
	assert.NoError(t, mock.ExpectationsWereMet(), "Not all expectations were met")
}
