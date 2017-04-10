package main

import (
	"github.com/kataras/iris"

	"github.com/fuyufjh/simple-gallery/obs/qiniu"
	"time"
)

func main() {

	obs := qiniu.New()

	iris.StaticCacheDuration = time.Duration(1)

	iris.Get("/list", func(ctx *iris.Context) {
		categories, err := obs.List()
		if err != nil {
			ctx.Error(err.Error(), iris.StatusInternalServerError)
			return
		}

		ctx.JSON(iris.StatusOK, categories)
	})

	//iris.Get("/", func(ctx *iris.Context) {
	//
	//})

	iris.Static("/static", "static", 1)

	iris.Listen("127.0.0.1:5700")
}
