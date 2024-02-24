package sql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"product-management/common"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConn *sql.DB
)

type dbConfig struct {
	host     string
	port     string
	username string
	password string
	database string
}

func newDBConfig() *dbConfig {
	return &dbConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_DATABASE"),
	}
}

func (c *dbConfig) GetHost() string {
	return c.host
}

func (c *dbConfig) GetPort() string {
	return c.port
}

func (c *dbConfig) GetUsername() string {
	return c.username
}

func (c *dbConfig) GetPassword() string {
	return c.password
}

func (c *dbConfig) GetDatabase() string {
	return c.database
}

func ConnectToDB() {
	var err error
	config := newDBConfig()
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=UTC&charset=utf8mb4",
		config.GetUsername(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetDatabase())
	DBConn, err = sql.Open(common.DB_DRIVER, mysqlInfo)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
}
