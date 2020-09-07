package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["company-manager/controllers:CompanyController"] = append(beego.GlobalControllerRouter["company-manager/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company-manager/controllers:CompanyController"] = append(beego.GlobalControllerRouter["company-manager/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company-manager/controllers:CompanyController"] = append(beego.GlobalControllerRouter["company-manager/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/deleteCompany`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company-manager/controllers:CompanyController"] = append(beego.GlobalControllerRouter["company-manager/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/getCompany`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company-manager/controllers:CompanyController"] = append(beego.GlobalControllerRouter["company-manager/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/update`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/GetAll`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"],
        beego.ControllerComments{
            Method: "GetAllInCompany",
            Router: `/GetAllInCompany`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/GetEmployee`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"],
        beego.ControllerComments{
            Method: "GetSliceByTime",
            Router: `/getSliceByTime`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["company-manager/controllers:EmployeeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/update`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
