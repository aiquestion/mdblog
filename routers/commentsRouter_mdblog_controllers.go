package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["mdblog/controllers:GithubController"] = append(beego.GlobalControllerRouter["mdblog/controllers:GithubController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

}
