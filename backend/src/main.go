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
	//ip:port/museumar/signin?code=xxxxxxxx
	app.Get("/museumar/signin", workflow.SignIn)

	authenticated := app.Party("/museumar/authenticated")
	authenticated.Use(workflow.GetJwtAuthenticator().Serve)
	//get museumhomepage data
	//ip:port/museumar/authenticated/museumhomepage
	authenticated.Get("/museumhomepage", workflow.MuseumHomePage)

	//listen and serve
	app.Run(iris.Addr(":8080"))
}
