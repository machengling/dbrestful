package models

import (
	"dbrestful/store"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type SelectParam struct {
	Params    map[string]SFieldParam
	Columns   map[string]SColumnParam
	TableName string   `json:"tablename,omitempty"`
	Limit     string   `json:"limit,omitempty"`
	Offset    string   `json:"offset,omitempty"`
	IsTotal   bool     `json:"istotal,omitempty"`
	GroupBy   []string `json:"groupby,omitempty"`
}

type SColumnParam struct {
	Alias string `json:"alias,omitempty"`
}

type SFieldParam struct {
	Value      interface{} `json:"value,omitempty"`
	IsOr       bool        `json:"is_or,omitempty"`
	FuzzyMatch bool        `json:"fuzzy_match,omitempty"`
	NotEqual   bool        `json:"not_equal,omitempty"`
}

type RowData struct {
	Rows      []orm.Params `json:"rows,omitempty"`
	TotalSize interface{}  `json:"total_size,omitempty"`
}

// Select 查找操作
func Select(param SelectParam) (rows RowData, err error) {
	o := store.GetDB()

	whereStr := whereSQL(param)
	columnStr := columnSQL(param)
	pageStr := pageSQL(param)
	groupbyStr := groupbySQL(param)
	sql := strings.Join([]string{"SELECT", columnStr, "FROM", param.TableName, whereStr, groupbyStr, pageStr}, " ")
	logs.Debug(sql)

	var maps []orm.Params
	_, err = o.Raw(sql).Values(&maps)
	rows.Rows = maps

	// 如果GetTotal为true ，则查询列表的数量
	if param.IsTotal == true {
		rows.TotalSize = SelectCount(param)
	}
	if err != nil {
		logs.Error(err)
		return rows, err
	}
	if err != nil {
		logs.Error(err)
	}
	return rows, err
}

// SelectCount 查询数量
func SelectCount(param SelectParam) (count interface{}) {
	o := store.GetDB()

	paramStr := whereSQL(param)
	sql := strings.Join([]string{"SELECT count(1) as count FROM", param.TableName, paramStr}, " ")
	logs.Debug(sql)

	var maps []orm.Params
	_, err := o.Raw(sql).Values(&maps)

	if err != nil {
		logs.Error(err)
		return nil
	}
	if err != nil {
		logs.Error(err)
		return nil
	}
	logs.Debug("SelectCount", maps, maps[0]["count"])
	return maps[0]["count"]
}

func whereSQL(param SelectParam) string {
	paramStr := ""
	// 拼接where语句
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
	return paramStr
}

func columnSQL(param SelectParam) string {
	columnStr := "*"
	// 拼接查询的列
	var columnList []string
	for k, v := range param.Columns {
		columnName := k
		if v.Alias != "" {
			columnName += " as " + v.Alias + " "
		}
		columnList = append(columnList, columnName)
	}
	if len(columnList) != 0 {
		columnStr = strings.Join(columnList, ",")
	}
	return columnStr
}

func pageSQL(param SelectParam) string {
	pageStr := ""
	if param.Limit != "" {
		pageStr = " LIMIT " + param.Limit

		if param.Offset != "" {
			pageStr += " OFFSET " + param.Offset
		}
	}

	return pageStr
}

func groupbySQL(param SelectParam) string {
	groupbyStr := ""
	if param.GroupBy != nil {
		groupbyStr = strings.Join(param.GroupBy, ",")
	}
	if groupbyStr != "" {
		groupbyStr = " GROUP BY " + groupbyStr + " "
	}
	return groupbyStr
}
