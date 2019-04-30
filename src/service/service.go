package service

import (
	"db/etcd"
	"db/pg"
	"github.com/kataras/iris"
	"router"
)

type Svc struct {
	iris.Application
	Router  *router.Router
	Dynamic *etcd.Client
	SysDB   *pg.SysDB
	AuthDB  *pg.AuthDB
}

func NewSvc() *Svc {
	return &Svc{}
}

func DefaultSvc() *Svc {
	svc := NewSvc()
	svc.Application = *iris.Default()
	svc.AuthDB = pg.DefaultAuthDB()
	svc.SysDB = pg.DefaultSysDB()
	svc.Dynamic = etcd.DefaultCli()
	svc.Router = router.NewRouter(&svc.Application)
	return svc
}
