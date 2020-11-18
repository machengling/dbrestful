package controllers

import (
	"dbrestful/models"
	"dbrestful/utils"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CreateTableController struct {
	beego.Controller
}

// GetRoot 定义根路由
func (s *CreateTableController) GetRoot() string {
	return "createtable"
}

// GetRouters 定义路由
func (s *CreateTableController) GetRouters() map[string]string {
	return map[string]string{
		"": "post:CreateTable",
	}
}

// @Title CreateTable
// @Description create users
// @Param	body		body 	models.CreateTableParam	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (c *CreateTableController) CreateTable() {
	param := models.CreateTableParam{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.Error(err)
		c.Ctx.ResponseWriter.WriteHeader(200)
		c.Data["json"] = utils.Response{Code: "200", Msg: "failed", Data: err}
		c.ServeJSON()
		return
	}

	err = models.Createtable(param)
	if err != nil {
		logs.Error(err)
		c.Ctx.ResponseWriter.WriteHeader(200)
		c.Data["json"] = utils.Response{Code: "200", Msg: "failed", Data: err}
		c.ServeJSON()
		return
	}
	c.Data["json"] = utils.Response{Code: "200", Msg: "success"}
	c.ServeJSON()
}
