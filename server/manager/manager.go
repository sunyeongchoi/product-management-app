package manager

import (
	"database/sql"
)

type ManagerService interface {
	SignUp(manager Manager) error
	Get(phone string) (Manager, error)
}

type DBManagerService struct {
	DBConn *sql.DB
}

func NewDBManagerService(dbConn *sql.DB) *DBManagerService {
	return &DBManagerService{
		DBConn: dbConn,
	}
}

func (s *DBManagerService) SignUp(manager Manager) error {
	query := "INSERT INTO manager (phone, password) values (?, ?)"
	_, err := s.DBConn.Exec(query, manager.Phone, manager.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *DBManagerService) Get(phone string) (Manager, error) {
	var manager Manager
	query := "SELECT id, phone, password FROM manager WHERE phone = ?"
	row := s.DBConn.QueryRow(query, phone)
	if err := row.Scan(&manager.ID, &manager.Phone, &manager.Password); err != nil {
		return Manager{}, err
	}
	return manager, nil
}