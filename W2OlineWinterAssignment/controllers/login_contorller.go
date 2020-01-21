package controllers

import (
	"W2OlineWinterAssignment/models"
	"W2OlineWinterAssignment/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post(){
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username",username,"password",password)

	id := models.QueryUserWithUnAndPwd(username,utils.MD5Hash(password))
	fmt.Println("id:",id)
	if id!=0 {
		this.Data["json"] = map[string]interface{}{"code":1,"message":"登录成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code":0,"message":"登录失败"}
	}
	this.ServeJSON()
}