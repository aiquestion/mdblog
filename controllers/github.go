package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"mdblog/models"
	"os/exec"
)

type GithubController struct {
	beego.Controller
}

// @router / [post]
func (o *GithubController) Post() {
	isValid := models.ValidateSignature(o.Ctx.Input.Header("X-Hub-Signature"), string(o.Ctx.Input.RequestBody))
	isPush := o.Ctx.Input.Header("X-GitHub-Event") == "push"
	// we only listen push event
	if isValid && isPush {
		var req GitHubRequest
		json.Unmarshal(o.Ctx.Input.RequestBody, &req)
		ref := "refs/heads/" + beego.AppConfig.String("githubbranch")
		if req.Ref == ref {
			var cmd = exec.Command("bash", "-c", "cd "+beego.AppConfig.String("localfolder")+"&& git pull")
			err := cmd.Run()
			fmt.Println(err)
		}
	}
	// just return
	o.ServeJSON()
}

// only need ref parameter to see if is master branch
type GitHubRequest struct {
	Ref string
}
