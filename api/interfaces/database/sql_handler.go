package database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewSQLHandler() *gorm.DB {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	db, err := gorm.Open("mysql", dbUser+":"+dbPass+"@tcp(db)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}
