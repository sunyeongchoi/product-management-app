package product

import (
	"errors"
	"net/http"
	"sync"

	"product-management/productmgm/common"
	"product-management/server"
	"product-management/server/product"
)

var (
	once          sync.Once
	productDBConn *product.DBProductService
)

func getProductDBConn() *product.DBProductService {
	once.Do(func() {
		productDBConn = product.NewDBProductService(server.DBConn)
	})
	return productDBConn
}

func Register(prod product.Product) (statusCode int, err error) {
	if prod.Size != common.SMALL && prod.Size != common.LARGE {
		return http.StatusInternalServerError, errors.New("잘못된 상품 사이즈 입니다.")
	}
	err = getProductDBConn().Register(prod)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func Update(id int, updateFields map[string]interface{}) (statusCode int, err error) {
	err = getProductDBConn().Update(id, updateFields)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func List(searchKeyword string, cursor int, limit int) (productList product.ProductList, statusCode int, err error) {
	productList, err = getProductDBConn().List(searchKeyword, cursor, limit)
	if err != nil {
		return product.ProductList{}, http.StatusInternalServerError, err
	}
	return productList, http.StatusOK, nil
}

func Get(id int) (prod product.Product, statusCode int, err error) {
	prod, err = getProductDBConn().Get(id)
	if err != nil {
		return prod, http.StatusInternalServerError, err
	}
	return prod, http.StatusOK, nil
}

func Delete(id int) (statusCode int, err error) {
	err = getProductDBConn().Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}