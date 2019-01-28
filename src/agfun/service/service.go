package service

import (
	"agfun/dbcentral/etcddb"
	"agfun/dbcentral/mysqldb"
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
	std = NewSvc(etcddb.GetCli(), mysqldb.GetSysDB(), mysqldb.GetAuthDB())
}
func GetDefaultSvc() *Svc {
	if std == nil {
		initStd()
	}
	return std
}
