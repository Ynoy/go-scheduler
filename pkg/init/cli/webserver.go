package cli

import (
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go-scheduler/pkg/web/router"
	"log"
	"sync"
	"time"
)

// 初始化Web Server UI，采用iris框架作为支持
func startInitWebServer() {
	app := iris.New()
	// 设置日志等级
	app.Logger().SetLevel("info")
	app.OnErrorCode(404, func(context iris.Context) {
		context.Redirect("/", iris.StatusMovedPermanently)
	})
	app.RegisterView(iris.HTML("./pages/dist", ".html").Binary(pages.Asset, pages.AssetNames))

	app.Get("/", func(context iris.Context) {
		if err := context.View("index.html"); err != nil {
			log.Println(err)
		}
	})
	mvc.Configure(app.Party("api/init"), router.Initialize)
	assertHandler := iris.StaticEmbeddedHandler("./pages/dist", pages.Asset, pages.AssetNames, false)
	app.SPA(assertHandler).AddIndexName("index.html")
	sg := new(sync.WaitGroup)
	defer sg.Wait()

	iris.RegisterOnInterrupt(func() {
		sg.Add(1)
		defer sg.Done()
		syncCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		if err := app.Shutdown(syncCtx); err != nil {
			log.Println(err)
		}
	})

	if err := app.Run(iris.Addr(":8701"), iris.WithOptimizations, iris.WithCharset("UTF-8"), iris.WithoutInterruptHandler); err != nil {
		log.Println(err)
	}

}