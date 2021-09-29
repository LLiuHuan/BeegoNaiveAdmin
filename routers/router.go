// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	v1 "BeegoNaiveAdmin/controllers/v1"
	"fmt"

	"github.com/beego/beego/v2/server/web/context"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&v1.LoginController{},
			),
		),
		//beego.NSNamespace("/user",
		//	beego.NSInclude(
		//		&controllers.UserController{},
		//	),
		//),
	)
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		fmt.Println(ctx.Request)
		fmt.Println(ctx.Request.URL)
	})
	beego.AddNamespace(ns)
}
