package models

import (
	"dbrestful/store"
	"fmt"
	"strings"

	"github.com/astaxie/beego/logs"
)

// CreateTableParam 创建数据表参数
type CreateTableParam struct {
	Params    map[string]CTFieldParam `json:"params,omitempty"`
	TableName string                  `json:"tablename,omitempty"`
}

// CTFieldParam 创建数据表字段参数
type CTFieldParam struct {
	FieldType string `json:"fieldtype,omitempty"`
	IsNotNull bool   `json:"isnotnull,omitempty"`
	IsUnique  bool   `json:"isunique,omitempty"`
	IsPK      bool   `json:"ispk,omitempty"`
}

func Createtable(createParam CreateTableParam) error {
	o := store.GetDB()
	logs.Debug(createParam)
	paramStr := ""
	for k, v := range createParam.Params {
		paramStr += fmt.Sprintf("\n %s %s ", k, v.FieldType)
		if v.IsNotNull == true {
			paramStr += " NOT NULL "
		}
		if v.IsPK == true {
			paramStr += " PRIMARY KEY "
		} else if v.IsUnique == true {
			paramStr += " UNIQUE "
		}
		paramStr += ","
	}
	paramStr = strings.TrimRight(paramStr, ",")
	sql := "\nCREATE TABLE " + createParam.TableName + " ( " + paramStr + " )"
	logs.Debug(sql)
	_, err := o.Raw(sql).Exec()
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
