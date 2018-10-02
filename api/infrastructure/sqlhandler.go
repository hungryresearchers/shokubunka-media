package infrastructure

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	conn, err := gorm.Open("mysql", dbUser+":"+dbPass+"@tcp(db)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Create(domain interface{}) error {
	err := handler.Conn.Create(domain).Error
	return err
}

func (handler *SqlHandler) Find(domain interface{}) error {
	err := handler.Conn.Where(domain).First(domain).Error
	return err
}
