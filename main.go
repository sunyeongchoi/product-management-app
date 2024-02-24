package main

import (
	mysql "database/sql"
	"log"
	"product-management/pkg/apiclient/manager"
	"product-management/pkg/apiclient/product"
	"product-management/productmgm/middleware"
	"product-management/server"

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
	// 로컬에서 돌릴 경우 주석 해제
	//err := godotenv.Load(filepath.Join("./", ".env"))
	//if err != nil {
	//	log.Fatal("Error loading .env file", err)
	//}
	r := setupRouter()
	server.ConnectToDB()

	defer func(dbConn *mysql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("error from close database connection")
		}
	}(server.DBConn)
	runErr := r.Run(":8080")
	if runErr != nil {
		log.Fatal("error from run 8080 port")
	}
}
