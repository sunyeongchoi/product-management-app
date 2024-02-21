package product

import (
	"database/sql"
	"errors"
	"product-management/models"
	mysql "product-management/sql"
	"time"
)

func Register(product models.Product) error {
	query := "INSERT INTO product (manager_id, category, price, `name`, description, `size`, expired_date) values (?, ?, ?, ?, ?, ?, ?)"
	_, err := mysql.DBConn.Exec(query, product.ManagerID, product.Category, product.Price, product.Name, product.Description, product.Size, product.ExpiredDate)
	if err != nil {
		return err
	}
	return nil
}

func Update(id int, updateFields map[string]interface{}) error {
	query := "UPDATE product SET "
	var params []interface{}
	var err error
	for key, value := range updateFields {
		query += key + "=?, "
		if key == "expired_date" {
			value, err = time.Parse(time.RFC3339, value.(string))
			if err != nil {
				return err
			}
		}
		params = append(params, value)
	}
	query = query[:len(query)-2] + " WHERE id = ?"
	params = append(params, id)
	_, err = mysql.DBConn.Exec(query, params...)
	return err
}

func List() ([]models.Product, error) {
	var products []models.Product
	query := "SELECT id, manager_id, category, price, name, description, size, expired_date FROM product"
	rows, err := mysql.DBConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// TODO: 페이지네이션 - cursor based pagination 기반으로, 1page 당 10개의 상품이 보이도록
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.ManagerID, &product.Category, &product.Price, &product.Name, &product.Description, &product.Size, &product.ExpiredDate); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func Get(id string) (models.Product, error) {
	var product models.Product
	query := "SELECT id, manager_id, category, price, name, description, size, expired_date FROM product WHERE id = ?"
	row := mysql.DBConn.QueryRow(query, id)
	if err := row.Scan(&product.ID, &product.ManagerID, &product.Category, &product.Price, &product.Name, &product.Description, &product.Size, &product.ExpiredDate); err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, errors.New("Product not found")
		}
		return models.Product{}, err
	}
	return product, nil
}

func Delete(id string) error {
	query := "DELETE FROM product WHERE id = ?"
	result, err := mysql.DBConn.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("No rows affected. Product not found.")
	}
	return nil
}

// TODO: 검색 (초성검색, 단어검색) - 예) 슈크림 라떼 → 검색가능한 키워드 : 슈크림, 크림, 라떼, ㅅㅋㄹ, ㄹㄸ
func Search() ([]models.Product, error) {
	var products []models.Product
	query := "SELECT id, manager_id, category, price, name, description, size, expired_date FROM product"
	rows, err := mysql.DBConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// TODO: 페이지네이션
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.ManagerID, &product.Category, &product.Price, &product.Name, &product.Description, &product.Size, &product.ExpiredDate); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
