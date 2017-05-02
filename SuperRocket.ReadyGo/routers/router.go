package routers

import (
	"ReadyGo/SuperRocket.ReadyGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
