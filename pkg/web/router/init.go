package router

import (
	"github.com/kataras/iris/mvc"
	"go-scheduler/pkg/web/controller/init"
)

func Initialize(app *mvc.Application) {
	app.Handle(new(init.Controller))
}
