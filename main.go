package main

import (
	mysql "database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"product-management/cmd/manager"
	"product-management/cmd/product"
	"product-management/middleware"
	"product-management/sql"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/signup", manager.GetManagerAPIManager().SignUp)
	r.POST("/login", manager.GetManagerAPIManager().Login)
	r.POST("/logout", manager.GetManagerAPIManager().LogOut)

	productGroup := r.Group("/management")
	productGroup.Use(middleware.TokenAuthMiddleware)
	{
		productGroup.POST("/product", product.GetProductAPIManager().Register)
		productGroup.PATCH("/product/:id", product.GetProductAPIManager().Update)
		productGroup.GET("/products", product.GetProductAPIManager().List)
		productGroup.GET("/product/:id", product.GetProductAPIManager().Get)
		productGroup.DELETE("/product/:id", product.GetProductAPIManager().Delete)
	}
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