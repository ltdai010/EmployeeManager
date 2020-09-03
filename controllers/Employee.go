package controllers

import (
	"company-manager/gen-go/company"
	"company-manager/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Employee
type EmployeeController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	company.Employee	true		"body for user content"
// @Success 200 {int} company.Employee.Id
// @Failure 403 body is empty
// @router / [post]
func (u *EmployeeController) Post() {
	var employee company.Employee
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &employee)
	uid := models.AddUser(employee)
	u.Data["json"] = map[string]string{"employeeID": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Param companyID query string true "company id"
// @Success 200 {object} company.Employee
// @router / [get]
func (u *EmployeeController) GetAll() {
	companyID := u.GetString("companyID")
	users := models.GetAllUsers(companyID)
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Param   companyID query string true         "The company ID"
// @Success 200 {object} company.Employee
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *EmployeeController) Get() {
	uid := u.GetString(":uid")
	companyID := u.GetString("companyID")
	if uid != "" {
		user, err := models.GetUser(uid, companyID)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		query 	string	true		"The uid you want to update"
// @Param	cid		query	string	true		"The cid"
// @Param	body	body	company.Employee		"The company information"
// @Success 200 {object} company.Employee
// @Failure 403 :something is empty
// @router /update [put]
func (u *EmployeeController) Put() {
	var e company.Employee
	uid := u.GetString("uid")
	cid := u.GetString("cid")
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &e)
	if uid != "" {
		err = models.UpdateUser(uid, e.Name, e.Address, int(e.Age), cid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = "done"
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Param   companyID query string true         "The company id"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *EmployeeController) Delete() {
	uid := u.GetString(":uid")
	companyID := u.GetString("companyID")
	models.DeleteUser(uid, companyID)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
