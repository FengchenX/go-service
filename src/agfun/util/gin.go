package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, gin.H{"Code": 0, "Msg": "success", "Data": resp})
}

func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"Code": -1, "Msg": err.Error(), "Data": nil})
}
