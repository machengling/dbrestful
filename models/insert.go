package models

import (
	"dbrestful/store"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/logs"
)

type InsertParam struct {
	Params    map[string]interface{} `json:"params,omitempty"`
	TableName string                 `json:"tablename,omitempty"`
}

// RowAffacted 影响的行数
type RowAffacted struct {
	Rows int64 `json:"row_affacted"`
}

// Insert 插入
func Insert(tablename string, param map[string]interface{}) (rows int64, err error) {
	paramStr := ""
	valueStr := ""
	for k, v := range param {
		// 判断传入的值是否为字符串
		t := reflect.TypeOf(v)
		paramStr += k + ","
		if t.Kind() == reflect.String {
			valueStr += "'" + fmt.Sprint(v) + "'" + ","
		} else {
			valueStr += fmt.Sprint(v) + ","
		}

	}
	paramStr = strings.TrimRight(paramStr, ",")
	valueStr = strings.TrimRight(valueStr, ",")
	o := store.GetDB()
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tablename, paramStr, valueStr)
	logs.Info(sql)
	res, err := o.Raw(sql).Exec()

	if err == nil {
		num, _ := res.RowsAffected()
		logs.Info("row affected nums:", num)
		return num, err
	}
	logs.Error(err)
	return 0, err

}
