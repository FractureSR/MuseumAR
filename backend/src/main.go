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
	//get the museum list
	//ip:port/museumar/allmuseum
	app.Get("/museumar/allmuseums", workflow.AllMuseums)

	authenticated := app.Party("/museumar/authenticated")
	authenticated.Use(workflow.GetJwtAuthenticator().Serve)
	//get museumhomepage data
	//ip:port/museumar/authenticated/museumhomepage?museum=xx&token=xx
	authenticated.Get("/museumhomepage", workflow.MuseumHomePage)
	//ip:port/museumar/authenticated/alltickets?museum=xx&token=xx
	authenticated.Get("/alltickets", workflow.AllTickets)
	//ip:port/museumar/authenticated/allsouvenirs?museum=xx&token=xx
	authenticated.Get("/allsouvenirs", workflow.AllSouvenirs)
	//ip:port/museumar/authenticated/scan?museum=xx&token=xx
	authenticated.Post("/scan", workflow.Scan)

	//listen and serve
	app.Run(iris.Addr(":8080"))
}
