package main

import (
	"github.com/astaxie/beego"
	_ "solo-ci/init"
	_ "solo-ci/routers"
)

func main() {
	beego.Run()
}