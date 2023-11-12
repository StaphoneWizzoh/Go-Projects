//package config
//
//import (
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jinzhu/gorm"
//)
//
//var (
//	db *gorm.DB
//)
//
//func Connect() {
//	dsn := "root:staphone@tcp(127.0.0.1:3306)/bookstore"
//	d, err := gorm.Open("mysql", dsn)
//	if err != nil {
//		panic(err)
//	}
//	db = d
//}
//
//func GetDB() *gorm.DB {
//	return db
//}

package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
