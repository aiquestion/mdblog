package main

import (
	_ "mdblog/docs"
	"mdblog/models"
	_ "mdblog/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	gitPath := beego.AppConfig.String("localfolder")
	go func() {
		stop := models.Watch(gitPath)
		<-stop
	}()

	beego.Run()
}
