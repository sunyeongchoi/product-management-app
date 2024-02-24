package main

import (
	mysql "database/sql"
	"log"

	"github.com/joho/godotenv"

	"product-management/cmd/manager"
	"product-management/cmd/product"
	"product-management/middleware"
	"product-management/sql"

	"github.com/gin-gonic/gin"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := setupRouter()
	sql.ConnectToDB()

	defer func(dbConn *mysql.DB) {
		err = dbConn.Close()
		if err != nil {
			log.Fatal("error from close database connection")
		}
	}(sql.DBConn)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("error from run 8080 port")
	}
}
