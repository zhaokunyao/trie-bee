package controllers

import (
	"github.com/astaxie/beego"
    "trie-bee/mytrie"
)

type SugController struct {
	beego.Controller
}

var o = map[string] interface{} {
    "error_code": 0,
    "error_msg": "ok",
    "data" : make([]string, 1),
}

func (this *SugController) Get() {
    s := this.GetString("prefix")

    if s == "" {
        o["error_code"]    = 1
        o["error_msg"]     = "empty prefix"
        this.Data["json"] = o
        this.ServeJSON()
        return
    }

    members := mytrie.PrefixMembersList(s)
    o["data"] = members
    beego.Info(o)
    this.Data["json"] =  o
    this.ServeJSON()

}

