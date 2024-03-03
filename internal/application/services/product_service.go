package services

import (
	"net/http"
	"product-management/internal/domain/entities"
	"product-management/internal/domain/repositories"
	"product-management/internal/interface/api/rest/request"
)

type ProductService struct {
	repository repositories.ProductRepository
}

func NewProductService(repository repositories.ProductRepository) *ProductService {
	return &ProductService{repository: repository}
}

func (s *ProductService) Register(productCommand *request.CreateProductRequest) (statusCode int, err error) {
	var newProduct = entities.NewProduct(0, productCommand.ManagerID, productCommand.Category, productCommand.Price, productCommand.Name, productCommand.Description, productCommand.Size, productCommand.ExpiredDate)
	validatedProduct, err := entities.NewValidatedProduct(newProduct)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = s.repository.Register(validatedProduct)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *ProductService) Update(id int, updateFields map[string]interface{}) (statusCode int, err error) {
	err = s.repository.Update(id, updateFields)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *ProductService) List(searchKeyword string, cursor int, limit int) (productList *entities.ProductList, statusCode int, err error) {
	productList, err = s.repository.List(searchKeyword, cursor, limit)
	if err != nil {
		return &entities.ProductList{}, http.StatusInternalServerError, err
	}
	return productList, http.StatusOK, nil
}

func (s *ProductService) Get(id int) (prod *entities.ValidatedProduct, statusCode int, err error) {
	prod, err = s.repository.Get(id)
	if err != nil {
		return prod, http.StatusInternalServerError, err
	}
	return prod, http.StatusOK, nil
}

func (s *ProductService) Delete(id int) (statusCode int, err error) {
	err = s.repository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}