package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"product-management/utils"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection() (*sql.DB, error) {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=UTC&charset=utf8mb4",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	return sql.Open(utils.DB_DRIVER, mysqlInfo)
}
