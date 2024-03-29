package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"product-management/internal/domain/entities"
	"product-management/internal/domain/repositories"
	"product-management/utils"
	"time"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) repositories.ProductRepository {
	return &ProductRepository{db: db}
}

func (s *ProductRepository) Register(product *entities.ValidatedProduct) error {
	// Map domain entity to DB model
	dbProduct := ToDBProduct(product)
	query := "INSERT INTO product (manager_id, category, price, `name`, description, `size`, expired_date) values (?, ?, ?, ?, ?, ?, ?)"
	_, err := s.db.Exec(query, dbProduct.ManagerID, dbProduct.Category, dbProduct.Price, dbProduct.Name, dbProduct.Description, dbProduct.Size, dbProduct.ExpiredDate)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductRepository) Update(id int, updateFields map[string]interface{}) error {
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
	_, err = s.db.Exec(query, params...)
	return err
}

func (s *ProductRepository) List(searchKeyword string, cursor int, limit int) (*entities.ProductList, error) {
	var products entities.ProductList
	var rows *sql.Rows
	var err error
	// 검색 (초성검색, 단어검색) - 예) 슈크림 라떼 → 검색가능한 키워드 : 슈크림, 크림, 라떼, ㅅㅋㄹ, ㄹㄸ
	if searchKeyword != "" {
		if utils.IsConsonants(searchKeyword) {
			// 초성검색 예) 슈크림 라떼 → 검색가능한 키워드 : ㅅㅋㄹ, ㄹㄸ
			whereStmt := utils.GetWhereClause(searchKeyword, "name")
			// 페이지네이션 - cursor based pagination 기반으로, 1page 당 기본 10개의 상품이 보이도록
			if cursor > 0 {
				// 첫번째 페이지가 아닐 경우
				query := fmt.Sprintf("SELECT id, manager_id, category, price, name, description, size, expired_date FROM product WHERE id < %d AND %s ORDER BY id DESC LIMIT %d", cursor, whereStmt, limit)
				rows, err = s.db.Query(query)
			} else {
				// 첫번째 페이지일 경우
				query := fmt.Sprintf("SELECT id, manager_id, category, price, name, description, size, expired_date FROM product WHERE %s ORDER BY id DESC LIMIT %d", whereStmt, limit)
				rows, err = s.db.Query(query)
			}
		} else {
			// 단어검색 예) 슈크림 라떼 → 검색가능한 키워드 : 슈크림, 크림, 라떼
			// 페이지네이션 - cursor based pagination 기반으로, 1page 당 기본 10개의 상품이 보이도록
			if cursor > 0 {
				// 첫번째 페이지가 아닐 경우
				query := fmt.Sprintf("SELECT id, manager_id, category, price, name, description, size, expired_date FROM product WHERE id < %d AND name LIKE '%%%s%%' ORDER BY id DESC LIMIT %d", cursor, searchKeyword, limit)
				rows, err = s.db.Query(query)
			} else {
				// 첫번째 페이지일 경우
				query := fmt.Sprintf("SELECT id, manager_id, category, price, name, description, size, expired_date FROM product WHERE name LIKE '%%%s%%' ORDER BY id DESC LIMIT %d", searchKeyword, limit)
				rows, err = s.db.Query(query)
			}
		}
	} else {
		// 페이지네이션 - cursor based pagination 기반으로, 1page 당 기본 10개의 상품이 보이도록
		if cursor > 0 {
			// 첫번째 페이지가 아닐 경우
			query := fmt.Sprintf("SELECT id, manager_id, category, price, name, description, size, expired_date FROM product WHERE id < %d ORDER BY id DESC LIMIT %d", cursor, limit)
			rows, err = s.db.Query(query)
		} else {
			// 첫번째 페이지일 경우
			query := fmt.Sprintf("SELECT id, manager_id, category, price, name, description, size, expired_date FROM product ORDER BY id DESC LIMIT %d", limit)
			rows, err = s.db.Query(query)
		}
	}
	if err != nil {
		return &entities.ProductList{}, err
	}

	defer rows.Close()
	if err = rows.Err(); err != nil {
		return &entities.ProductList{}, err
	}
	for rows.Next() {
		var product entities.Product
		if err = rows.Scan(&product.ID, &product.ManagerID, &product.Category, &product.Price, &product.Name, &product.Description, &product.Size, &product.ExpiredDate); err != nil {
			return &entities.ProductList{}, err
		}
		products.Items = append(products.Items, product)
	}

	if len(products.Items) > 0 {
		products.Metadata.Cursor = products.Items[len(products.Items)-1].ID
	}
	return &products, nil
}

func (s *ProductRepository) Get(id int) (*entities.ValidatedProduct, error) {
	var product Product
	query := "SELECT id, manager_id, category, price, name, description, size, expired_date FROM product WHERE id = ?"
	row := s.db.QueryRow(query, id)
	if err := row.Scan(&product.ID, &product.ManagerID, &product.Category, &product.Price, &product.Name, &product.Description, &product.Size, &product.ExpiredDate); err != nil {
		if err == sql.ErrNoRows {
			return &entities.ValidatedProduct{}, errors.New("Product not found")
		}
		return &entities.ValidatedProduct{}, err
	}
	return FromDBProduct(&product)
}

func (s *ProductRepository) Delete(id int) error {
	query := "DELETE FROM product WHERE id = ?"
	result, err := s.db.Exec(query, id)
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
