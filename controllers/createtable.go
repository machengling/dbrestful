package controllers

import (
	"dbrestful/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CreateTableController struct {
	beego.Controller
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
		c.Data["json"] = Response{Code: "200", Msg: "failed", Data: err}
		c.ServeJSON()
		return
	}

	err = models.Createtable(param)
	if err != nil {
		logs.Error(err)
		c.Ctx.ResponseWriter.WriteHeader(200)
		c.Data["json"] = Response{Code: "200", Msg: "failed", Data: err}
		c.ServeJSON()
		return
	}
	c.Data["json"] = Response{Code: "200", Msg: "success"}
	c.ServeJSON()
}
