package router

import (
	"agfun/controller"
	"agfun/dbcentral/etcddb"
	"agfun/dbcentral/mysqldb"
	"agfun/jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, []string{"session", "accept"}...)
	router.Use(cors.New(config), jwt.AuthMiddleWare(pg.GetAuthDB(), etcddb.GetCli()))

	router.GET("/", controller.Hello)
	initFreeVideoRouter(router)
	return router
}
