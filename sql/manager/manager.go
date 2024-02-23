package manager

import (
	"database/sql"
	"product-management/models"
)

type ManagerService interface {
	SignUp(manager models.Manager) error
	Get(phone string) (models.Manager, error)
}

type DBManagerService struct {
	DBConn *sql.DB
}

func NewDBManagerService(dbConn *sql.DB) *DBManagerService {
	return &DBManagerService{
		DBConn: dbConn,
	}
}

func (s *DBManagerService) SignUp(manager models.Manager) error {
	query := "INSERT INTO manager (phone, password) values (?, ?)"
	_, err := s.DBConn.Exec(query, manager.Phone, manager.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *DBManagerService) Get(phone string) (models.Manager, error) {
	var manager models.Manager
	query := "SELECT id, phone, password FROM manager WHERE phone = ?"
	row := s.DBConn.QueryRow(query, phone)
	if err := row.Scan(&manager.ID, &manager.Phone, &manager.Password); err != nil {
		return models.Manager{}, err
	}
	return manager, nil
}