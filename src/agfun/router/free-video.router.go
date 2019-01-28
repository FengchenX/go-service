package router

import (
	"agfun/controller"
	"github.com/gin-gonic/gin"
)

func initFreeVideoRouter(router *gin.Engine) {
	group := router.Group("/free-video")
	group.GET("/", controller.GetFreeVideos)
	group.POST("/", controller.AddFreeVideos)
	group.PUT("/{id}", controller.UpdateFreeVideo)
	group.DELETE("/{id}", controller.DelFreeVideo)
}
