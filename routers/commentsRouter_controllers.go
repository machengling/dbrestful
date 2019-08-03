package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["dbrestful/controllers:InsertController"] = append(beego.GlobalControllerRouter["dbrestful/controllers:InsertController"],
        beego.ControllerComments{
            Method: "Insert",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
