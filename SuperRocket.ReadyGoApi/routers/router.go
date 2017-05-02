package routers

import (
	"ReadyGo/SuperRocket.ReadyGoApi/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
