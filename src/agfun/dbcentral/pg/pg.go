package pg

import (
	"agfun/conf"
	"agfun/entity"
	"agfun/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	*gorm.DB
}

func NewDB() *DB {
	return &DB{}
}
func (db *DB) initDB(a ...interface{}) {
	d, e := gorm.Open("postgres", conf.AgfunInst().SysDB)
	if e != nil {
		log.Fatal(e)
	}
	db.DB = d
	db.addTable()
}
func (db *DB) addTable(a ...interface{}) {
	if migrate := db.DB.AutoMigrate(a...); migrate.Error != nil {
		log.Fatal(migrate.Error)
	}
}
type SysDB struct {
	DB
}

func SysCreateTable() {
	if db := sysdb.AutoMigrate(&entity.Video{}); db.Error != nil {
		panic(db.Error)
	}
}

//&parseTime=true&loc=Local
func initSysDB() {
	db, err := gorm.Open("postgres",
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
	db, err := gorm.Open("postgres",
		conf.AgfunInst().AuthDB)
	if err != nil {
		log.Fatal(err)
	}
	authdb = db
	authdb.LogMode(true)
}
