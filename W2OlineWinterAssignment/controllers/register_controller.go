package controllers

import (
	"W2OlineWinterAssignment/models"
	"W2OlineWinterAssignment/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/labstack/gommon/log"
)

type RegisterController struct{
	beego.Controller
}

//GET方法
func (this *RegisterController) Get(){
	this.TplName = "register.html"
}

//POST方法
func (this *RegisterController) Post(){
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Print(username,password,repassword)
	log.Info(username,password,repassword)

	//判断该用户名是否被注册过 用户名是唯一的
	id := models.QueryUserWithUsername(username)
	fmt.Println(id)
	if id != 0{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"用户名已存在"}
		this.ServeJSON()
		return
	}

	//密码采用MD5哈希方式加密，在登录时也是把用户的密码MD5哈希后进行比较
	password = utils.MD5Hash(password)
	fmt.Println("MD5Hash:",password)

	user := models.User{0,username,password,0}
	_, err := models.InsertUser(user)
	if err != nil{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"注册失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code":1,"message":"注册成功"}
	}
	this.ServeJSON()
}

