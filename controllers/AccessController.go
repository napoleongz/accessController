package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"accessController/utils"
	"accessController/models"
)

type AccessController struct {
	BaseController
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *AccessController) GetAccess(){
	login := Login{}
	requestBody := c.Ctx.Input.RequestBody
	err := json.Unmarshal(requestBody, &login)
	if err != nil {
		beego.Error(err)
	}
	pw := login.Password
	un := login.Username
	pw = utils.String2md5(pw)
	user, err := models.LoginByUserName(un, pw)
	beego.Info(user)
	c.setTpl("accesscontroller/getaccess.html")
}