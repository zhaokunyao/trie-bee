package main

import (
	_ "trie-bee/routers"
	"github.com/astaxie/beego"
    "fmt"
    "trie-bee/mytrie"
)

func main() {
    mytrie.Init()
    fmt.Println("main.go")
	beego.Run()
}

