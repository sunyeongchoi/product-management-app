package mysql

import (
	"database/sql"
	"product-management/internal/domain/entities"
	"product-management/internal/domain/repositories"
)

type ManagerRepository struct {
	db *sql.DB
}

func NewManagerRepository(db *sql.DB) repositories.ManagerRepository {
	return &ManagerRepository{db: db}
}

func (s *ManagerRepository) SignUp(manager *entities.ValidatedManager) error {
	// Map domain entity to DB model
	dbManager := ToDBManager(manager)
	query := "INSERT INTO manager (phone, password) values (?, ?)"
	_, err := s.db.Exec(query, dbManager.Phone, dbManager.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *ManagerRepository) Get(phone string) (*entities.ValidatedManager, error) {
	var manager Manager
	query := "SELECT id, phone, password FROM manager WHERE phone = ?"
	row := s.db.QueryRow(query, phone)
	if err := row.Scan(&manager.ID, &manager.Phone, &manager.Password); err != nil {
		return &entities.ValidatedManager{}, err
	}
	// Map back to domain entity
	return FromDBManager(&manager)
}
