package util

import "github.com/kataras/iris"

func Success(c iris.Context, data interface{}) {
	c.JSON(iris.Map{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}
func Fail(c iris.Context, err error) {
	c.JSON(iris.Map{
		"code": -1,
		"msg":  err.Error(),
		"data": nil,
	})
}
