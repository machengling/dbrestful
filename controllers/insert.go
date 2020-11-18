package controllers

import (
	"dbrestful/models"
	"dbrestful/utils"
	"encoding/json"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

// InsertController ...
type InsertController struct {
	beego.Controller
}

// GetRoot 定义根路由
func (s *InsertController) GetRoot() string {
	return "insert"
}

// GetRouters 定义路由
func (s *InsertController) GetRouters() map[string]string {
	return map[string]string{
		"": "post:Insert",
	}
}

// @Title Insert ...
// @Description create users
// @Param	body		body 	models.InsertParam	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *InsertController) Insert() {
	param := models.InsertParam{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.Error(err)
		u.Ctx.ResponseWriter.WriteHeader(200)
		u.Data["json"] = utils.Response{Code: "200", Msg: "failed", Data: err}
		u.ServeJSON()
		return
	}

	rows, err := models.Insert(param.TableName, param.Params)
	if err != nil {
		logs.Error(err)
		u.Ctx.ResponseWriter.WriteHeader(200)
		u.Data["json"] = utils.Response{Code: "200", Msg: "failed", Data: err}
		u.ServeJSON()
		return
	}
	u.Data["json"] = utils.Response{Code: "200", Msg: "success", Data: models.RowAffacted{Rows: rows}}
	u.ServeJSON()
}
