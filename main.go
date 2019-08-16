package main

import (
	_ "accessController/routers"
	_ "accessController/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
