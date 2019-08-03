package controllers

import (
	"dbrestful/models"
	"encoding/json"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

// Operations about Users
type InsertController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.InsertParam	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /insert [post]
func (u *InsertController) Insert() {
	param := models.InsertParam{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &param)
	if err != nil {
		logs.Error(err)
		u.Data["json"] = map[string]string{"error": err.Error()}
		return
	}

	rows, err := models.Insert(param.TableName, param.Param)
	if err != nil {
		logs.Error(err)
		u.Data["json"] = map[string]string{"error": err.Error()}
		return
	}
	u.Data["json"] = models.RowAffacted{Rows: rows}
	u.ServeJSON()
}
