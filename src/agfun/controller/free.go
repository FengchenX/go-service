package controller

import (
	"agfun/dto"
	"agfun/entity"
	"agfun/service"
	"agfun/util"
	"github.com/gin-gonic/gin"
)

func GetFreeVideos(c *gin.Context) {
	req, e := decodeGetFreeVideos(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	videos, e := service.GetDefaultSvc().GetFreeVideos(req)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, videos)
}

func decodeGetFreeVideos(c *gin.Context) (dto.GetVideos, error) {
	var req dto.GetVideos
	filter, e := util.ParsePageFilter(c)
	if e != nil {
		return req, e
	}
	token := c.GetHeader("session")
	req.Filter = filter
	req.Token = token
	return req, nil
}

func AddFreeVideos(c *gin.Context) {
	videos, e := decodeCreateFreeVideos(c)
	if e != nil {
		util.Fail(c, e)
		return
	}
	e = service.GetDefaultSvc().AddFreeVideos(videos)
	if e != nil {
		util.Fail(c, e)
		return
	}
	util.Success(c, videos)
}

func decodeCreateFreeVideos(c *gin.Context) ([]*entity.FreeVideo, error) {
	var req []*entity.FreeVideo
	e := c.BindJSON(&req)
	return req, e
}

func UpdateFreeVideo(c *gin.Context) {

}

func DelFreeVideo(c *gin.Context) {

}
