package pg

import (
	"agfun/conf"
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
func (db *DB) ConnectDB(addr string) {
	d, e := gorm.Open("postgres", addr)
	if e != nil {
		log.Fatal(e)
	}
	db.DB = d
	db.LogMode(true)
}


type SysDB struct {
	DB
}
func NewSysDB() *SysDB {
	return &SysDB{}
}
func DefaultSysDB() *SysDB {
	sys:=NewSysDB()
	sys.ConnectDB(conf.AgfunInst().SysDB)
	return sys
}
type AuthDB struct {
	DB
}
func NewAuthDB() *AuthDB {
	return &AuthDB{}
}
func DefaultAuthDB() *AuthDB{
	auth:=NewAuthDB()
	auth.ConnectDB(conf.AgfunInst().AuthDB)
	return auth
}
