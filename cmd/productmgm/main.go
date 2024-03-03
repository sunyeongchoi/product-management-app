package main

import (
	"log"
	"os"
	"path/filepath"
	"product-management/internal/application/services"
	"product-management/internal/infrastructure/db/mysql"
	"product-management/internal/interface/api/rest"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("IS_PRODUCTION") == "" {
		err := godotenv.Load(filepath.Join("./", ".env"))
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	sqlDB, err := mysql.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	managerRepo := mysql.NewManagerRepository(sqlDB)
	productRepo := mysql.NewProductRepository(sqlDB)

	managerService := services.NewManagerService(managerRepo)
	productService := services.NewProductService(productRepo)

	r := gin.Default()
	rest.NewManagerController(r, managerService)
	rest.NewProductController(r, productService)

	runErr := r.Run(":8080")
	if runErr != nil {
		log.Fatal("error from run 8080 port")
	}
}
