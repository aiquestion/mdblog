package controllers

import (
	"github.com/astaxie/beego"
	"mdblog/models"
)

type GithubController struct {
	beego.Controller
}

// @router / [post]
func (o *GithubController) Post() {
	isValid := models.ValidateSignature(o.Ctx.Input.Header("X-Hub-Signature"), string(o.Ctx.Input.RequestBody))
	o.Data["json"] = map[string]bool{"ObjectId": isValid}
	o.ServeJSON()
}
