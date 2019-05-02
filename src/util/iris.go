package util

import "github.com/kataras/iris"

func Success(c iris.Context, data interface{}) {
	c.JSON(iris.Map{
		"Code": 0,
		"Msg":  "success",
		"Data": data,
	})
}
func Fail(c iris.Context, err error) {
	c.JSON(iris.Map{
		"Code": -1,
		"Msg":  "fail",
		"Data": err,
	})
}
