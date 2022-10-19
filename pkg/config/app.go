package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() *gorm.DB {
	//user:password@tcp(0.0.0.0:3306)/table?charset=utf8mb4&parseTime=True&loc=Local
	connStr := fmt.Sprintf("%s/bookstore?charset=utf8mb4&parseTime=True&loc=Local",os.Getenv("mysql") )
	d, err := gorm.Open("mysql", connStr)
	if (err != nil){
		panic(err)
	}

	db = d
	return db
}

func GetDB() *gorm.DB {
	return db
}
