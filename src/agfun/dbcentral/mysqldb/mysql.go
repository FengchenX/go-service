package mysqldb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

//&parseTime=true&loc=Local
func initSysDB() {
	db, err := gorm.Open("mysql",
		"root:feng@tcp(localhost:3306)/agfun?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	sysdb = db
}

var sysdb *gorm.DB

func GetSysDB() *gorm.DB {
	if sysdb == nil {
		initSysDB()
	}
	return sysdb
}

var authdb *gorm.DB

func GetAuthDB() *gorm.DB {
	if authdb == nil {
		initAuthDB()
	}
	return authdb
}

func initAuthDB() {
	db, err := gorm.Open("mysql",
		"root:feng@tcp(localhost:3306)/auth?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	authdb = db
	authdb.LogMode(true)
}
