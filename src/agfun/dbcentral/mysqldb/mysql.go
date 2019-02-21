package mysqldb

import (
	"agfun/conf"
	"agfun/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func SysCreateTable() {
	if db := sysdb.AutoMigrate(&entity.Video{}); db.Error != nil {
		panic(db.Error)
	}
}

//&parseTime=true&loc=Local
func initSysDB() {
	db, err := gorm.Open("mysql",
		conf.AgfunInst().SysDB)
	if err != nil {
		log.Fatal(err)
	}
	sysdb = db
	SysCreateTable()
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
		conf.AgfunInst().AuthDB)
	if err != nil {
		log.Fatal(err)
	}
	authdb = db
	authdb.LogMode(true)
}
