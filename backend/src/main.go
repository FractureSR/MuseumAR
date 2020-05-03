package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"

	//"github.com/kataras/iris/v12/middleware/recover"

	"MuseumAR_Backend/workflow"
)

func main() {
	//construct the app
	app := iris.New()
	//config middlewares
	app.Use(logger.New())
	//app.Use(recover.New())

	//mux
	//receive code to authorize and create a session
	app.Get("/signin", workflow.SignIn)

	//listen and serve
	app.Run(iris.Addr(":8080"))
}
