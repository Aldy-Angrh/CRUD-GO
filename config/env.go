package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DBInit method to init dv
func DBInit() *gorm.DB {

	dbdriver := os.Getenv("DB_DRIVER")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")

	db, err := gorm.Open(dbdriver, dbuser+":"+dbpass+"@("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8&parseTime=false&loc=Asia%2FJakarta")
	if err != nil {
		panic(err)
	}
	return db
}
