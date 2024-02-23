package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	DBConn *sql.DB
)

// TODO: 설정 값으로 빼기
const (
	host     = "localhost"
	port     = 3306
	username = "admin"
	password = "passwd"
	database = "productmgm"
)

type dbConfig struct {
	host     string
	port     int
	username string
	password string
	database string
}

func newDBConfig() *dbConfig {
	return &dbConfig{
		host:     host,
		port:     port,
		username: username,
		password: password,
		database: database,
	}
}

func (c *dbConfig) GetHost() string {
	return c.host
}

func (c *dbConfig) GetPort() int {
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
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=UTC&charset=utf8mb4",
		config.GetUsername(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetDatabase())
	DBConn, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
}
