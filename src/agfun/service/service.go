package service

import (
	"agfun/db/etcd"
	"agfun/db/pg"
	"github.com/kataras/iris"
)

type Svc struct {
	iris.Application
	Dynamic *etcd.Client
	SysDB   *pg.SysDB
	AuthDB  *pg.AuthDB
}

func NewSvc() *Svc {
	return &Svc{}
}

func DefaultSvc() *Svc{
	svc:=NewSvc()
	svc.Application = *iris.Default()
	svc.AuthDB = pg.DefaultAuthDB()
	svc.SysDB = pg.DefaultSysDB()
	//svc.Dynamic =
	return svc
}
