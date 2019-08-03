package store

import (
	"fmt"
	"sync"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func init() {
	dbuser := beego.AppConfig.String("store::dbuser")
	if dbuser == "" {
		panic("数据库dbuser未设置")
	}
	dbpwd := beego.AppConfig.String("store::dbpwd")
	if dbpwd == "" {
		panic("数据库dbpwd未设置")
	}
	dburl := beego.AppConfig.String("store::dburl")
	if dburl == "" {
		panic("数据库dburl未设置")
	}
	dbport := beego.AppConfig.String("store::dbport")
	if dbport == "" {
		panic("数据库dbport未设置")
	}
	dbsheet := beego.AppConfig.String("store::dbsheet")
	if dbsheet == "" {
		panic("数据库dbsheet未设置")
	}
	dbtype := beego.AppConfig.String("store::dbtype")
	if dbtype == "" {
		panic("数据库dbtype未设置")
	}
	switch dbtype {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		err := orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpwd, dburl, dbport, dbsheet))
		if err != nil {
			panic(err)
		}
	case "postgres":
		orm.RegisterDriver("postgres", orm.DRPostgres)
		err := orm.RegisterDataBase("postgres", "postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disabled", dbuser, dbpwd, dburl, dbport, dbsheet))
		if err != nil {
			panic(err)
		}
	}
}

// GetDB 从连接池中获取数据库实例对象
func GetDB() orm.Ormer {
	p := sync.Pool{
		New: func() interface{} {
			return orm.NewOrm()
		},
	}
	return p.Get().(orm.Ormer)
}
