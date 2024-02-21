package main

import (
	mysql "database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"product-management/cmd/product"
	"product-management/sql"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/product", product.GetProductAPIManager().Register)
	r.PATCH("/product/:id", product.GetProductAPIManager().Update)
	r.GET("/products", product.GetProductAPIManager().List)
	r.GET("/product/:id", product.GetProductAPIManager().Get)
	r.DELETE("/product/:id", product.GetProductAPIManager().Delete)
	return r
}

func main() {
	r := setupRouter()
	sql.ConnectToDB()
	defer func(dbConn *mysql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("error from close database connection")
		}
	}(sql.DBConn)
	r.Run(":8080")
}