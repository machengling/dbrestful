package controllers

import (
	"dbrestful/models"
	"dbrestful/utils"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type SelectController struct {
	beego.Controller
}

// GetRoot 定义根路由
func (s *SelectController) GetRoot() string {
	return "select"
}

// GetRouters 定义路由
func (s *SelectController) GetRouters() map[string]string {
	return map[string]string{
		"": "post:Select",
	}
}

// @Title Select
// @Description create users
// @Param	body		body 	models.CreateTableParam	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (d *SelectController) Select() {
	param := models.SelectParam{}
	err := json.Unmarshal(d.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.Error(err)
		d.Ctx.ResponseWriter.WriteHeader(200)
		d.Data["json"] = utils.Response{Code: "200", Msg: "failed", Data: err}
		d.ServeJSON()
		return
	}

	rows, err := models.Select(param)
	if err != nil {
		logs.Error(err)
		d.Ctx.ResponseWriter.WriteHeader(200)
		d.Data["json"] = utils.Response{Code: "200", Msg: "failed", Data: err}
		d.ServeJSON()
		return
	}
	d.Data["json"] = utils.Response{Code: "200", Msg: "success", Data: rows}
	d.ServeJSON()
}
