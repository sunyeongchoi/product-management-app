package manager

import (
	"database/sql"
	"errors"
	"product-management/models"
	mysql "product-management/sql"
)

func SignUp(manager models.Manager) error {
	query := "INSERT INTO manager (phone, password) values (?, ?)"
	_, err := mysql.DBConn.Exec(query, manager.Phone, manager.Password)
	if err != nil {
		return err
	}
	return nil
}

func Get(phone string) (models.Manager, error) {
	var manager models.Manager
	query := "SELECT id, phone, password FROM manager WHERE phone = ?"
	row := mysql.DBConn.QueryRow(query, phone)
	if err := row.Scan(&manager.ID, &manager.Phone, &manager.Password); err != nil {
		if err == sql.ErrNoRows {
			return models.Manager{}, errors.New("Manager not found")
		}
		return models.Manager{}, err
	}
	return manager, nil
}