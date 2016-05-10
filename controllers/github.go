package controllers

import (
	"github.com/astaxie/beego"
)

type GithubController struct {
	beego.Controller
}

// @router / [post]
func (o *GithubController) Post() {
	o.Data["json"] = map[string]string{"ObjectId": "1"}
	o.ServeJSON()
}
