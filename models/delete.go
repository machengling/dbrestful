package models

import (
	"dbrestful/store"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/logs"
)

type DeleteParam struct {
	Params    map[string]DFieldParam
	TableName string
}

type DFieldParam struct {
	Value      interface{} `json:"value,omitempty"`
	IsOr       bool        `json:"is_or,omitempty"`
	FuzzyMatch bool        `json:"fuzzy_match,omitempty"`
	NotEqual   bool        `json:"not_equal,omitempty"`
}

// Delete 删除操作
func Delete(param DeleteParam) (int64, error) {
	o := store.GetDB()

	paramStr := ""
	for k, v := range param.Params {
		// 判断参数或与
		if v.IsOr == true {
			paramStr += " OR "
		} else {
			paramStr += " AND "
		}
		// 判断参数类型
		value := fmt.Sprint(v.Value)
		t := reflect.TypeOf(v.Value)
		if t.Kind() == reflect.String {
			value = "'" + fmt.Sprint(v.Value) + "'"
		}
		// 判断参数是like，not null
		if v.FuzzyMatch == true {
			if v.NotEqual == true {
				paramStr += k + " NOT LIKE " + value
			} else {
				paramStr += k + " LIKE " + value
			}
		} else {
			if v.NotEqual == true {
				paramStr += k + " != " + value
			} else {
				paramStr += k + " = " + value
			}
		}
	}
	// 去除多余的前缀OR/AND
	paramStr = strings.TrimLeft(paramStr, " OR")
	paramStr = strings.TrimLeft(paramStr, " AND")
	// 添加前缀where
	if paramStr != "" {
		paramStr = " WHERE " + paramStr
	}
	sql := "DELETE FROM " + param.TableName + paramStr
	logs.Debug(sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		logs.Error(err)
	}
	return rows, err
}
