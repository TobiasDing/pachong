package main

import (
	_ "pachong/routers"
	beego "github.com/astaxie/beego/server/web"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

