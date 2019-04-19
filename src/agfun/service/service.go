package service

import (
	"agfun/dbcentral/etcddb"
	"agfun/dbcentral/mysqldb"
	"agfun/dbcentral/pg"
	"github.com/jinzhu/gorm"
)

type Svc struct {
	Dynamic *etcddb.Client
	SysDB   *gorm.DB
	AuthDB  *gorm.DB
}

func NewSvc(dynamic *etcddb.Client, sys, auth *gorm.DB) *Svc {
	return &Svc{
		Dynamic: dynamic,
		SysDB:   sys,
		AuthDB:  auth,
	}
}

var std *Svc

func initStd() {
	std = NewSvc(etcddb.GetCli(), pg.GetSysDB(), pg.GetAuthDB())
}
func GetDefaultSvc() *Svc {
	if std == nil {
		initStd()
	}
	return std
}
