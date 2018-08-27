package infrastructure

import (
	"os"
	"shokubunka-media/api/src/interfaces/database"

	"github.com/jinzhu/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	conn, err := gorm.Open("mysql", dbUser+":"+dbPass+"@tcp(hungry_db)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Create(entity *interface{}) (database.Result, error) {
	err := handler.Conn.Create(entity).Error
	return entity, err
}
