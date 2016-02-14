package routers

import (
	"trie-bee/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/sug", &controllers.SugController{})
}
