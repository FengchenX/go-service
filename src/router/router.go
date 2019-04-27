package router
//
//import (
//	"controller"
//	"db/etcd"
//	"db/mysqldb"
//	"jwt"
//	"github.com/gin-contrib/cors"
//	"github.com/gin-gonic/gin"
//)
//
//func Init() *gin.Engine {
//	router := gin.Default()
//	config := cors.DefaultConfig()
//	config.AllowAllOrigins = true
//	config.AllowHeaders = append(config.AllowHeaders, []string{"session", "accept"}...)
//	router.Use(cors.New(config), jwt.AuthMiddleWare(pg.GetAuthDB(), etcd.GetCli()))
//
//	router.GET("/", controller.Hello)
//	initFreeVideoRouter(router)
//	return router
//}
