package routers

import (
	"W2OlineWinterAssignment/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/register",&controllers.RegisterController{})
    beego.Router("/login",&controllers.LoginController{})
}
