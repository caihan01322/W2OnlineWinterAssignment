package main

import (
	_ "W2OlineWinterAssignment/routers"
	"W2OlineWinterAssignment/utils"
	"github.com/astaxie/beego"
)

func main() {
	utils.InitMysql()
	beego.Run()
}

