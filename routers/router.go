package routers

import (
	beego "github.com/astaxie/beego/server/web"
	"pachong/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:Post")
	beego.Router("/craw_movie", &controllers.CrawMovieController{}, "*:CrawMovie")

}
