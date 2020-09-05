package controllers

import (
	"company-manager/gen-go/company"
	"company-manager/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

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
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &employee)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		uid := models.AddEmployee(employee)
		u.Data["json"] = map[string]string{"employeeID": uid}
	}
	u.ServeJSON()
}

// @Title GetAllInCompany
// @Description get all Users
// @Param companyID query string true "company id"
// @Success 200 {object} company.Employee
// @router /GetAllInCompany [get]
func (u *EmployeeController) GetAllInCompany() {
	companyID := u.GetString("companyID")
	users := models.GetAllCompanyEmployee(companyID)
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} company.Employee
// @router /GetAll [get]
func (u *EmployeeController) GetAll() {
	users, err := models.GetAllEmployee()
	if err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = users
	}
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
		user, err := models.GetEmployee(uid, companyID)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title GetSlice
// @Description get user by uid
// @Param   companyID query string true         "The company ID"
// @Param   start query int true         "Start point"
// @Param   count query int true         "Number of elements"
// @Success 200 {object} company.Employee
// @Failure 403 :uid is empty
// @router /getSlice [get]
func (u *EmployeeController) GetSlice() {
	companyID := u.GetString("companyID")
	start, err := u.GetInt("start")
	if err != nil {
		u.Data["json"] = "error"
	}
	count, err := u.GetInt("count")
	if err == nil {
		list, err := models.GetEmployeeSlice(companyID, start, count)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = list
		}
	}
	u.ServeJSON()
}

// @Title GetSliceByTime
// @Description get user by uid
// @Param   companyID query string true         "The company ID"
// @Param   first query string true         "From this day(dd-mm-yyyy)"
// @Param   last query string true         "To this day (dd-mm-yyyy)"
// @Success 200 {object} company.Employee
// @Failure 403 :uid is empty
// @router /getSliceByTime [get]
func (u *EmployeeController) GetSliceByTime() {
	first := u.GetString("first")
	last := u.GetString("last")
	companyID := u.GetString("companyID")
	firstDay, err := stringToDate(first)
	if err != nil {
		u.Data["json"] = err.Error()
	}
	lastDay, err := stringToDate(last)
	if err != nil {
		u.Data["json"] = err.Error()
	}
	if err == nil {
		list, err := models.GetEmployeeSliceInTime(companyID, firstDay, lastDay)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = list
		}
	}
	u.ServeJSON()
}


func stringToDate(date string) (*company.Date, error) {
	var day []string
	day = strings.Split(date, "-")
	if len(day) != 3 {
		return nil, errors.New("wrong type")
	}
	dd, err := strconv.Atoi(day[0])
	if err != nil {
		return nil, err
	}
	mm, err := strconv.Atoi(day[1])
	if err != nil {
		return nil, err
	}
	yyyy, err := strconv.Atoi(day[2])
	if err != nil {
		return nil, err
	}
	d := company.Date{Day: company.Int(dd), Month: company.Int(mm), Year: company.Int(yyyy)}
	return &d, nil
}

// @Title Update
// @Description update the user
// @Param	uid		query 	string	true		"The uid you want to update"
// @Param	cid		query	string	true		"The cid"
// @Param	body	body	templateType.updateEmployeeForm		"The company information"
// @Success 200 {object} company.Employee
// @Failure 403 :something is empty
// @router /update [put]
func (u *EmployeeController) Put() {
	var e company.Employee
	uid := u.GetString("uid")
	cid := u.GetString("cid")
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &e)
	if uid != "" {
		err = models.UpdateEmployee(uid, e.GetName(), e.GetAddress(), e.GetDate(), cid)
		if err != nil {
			u.Data["json"] = err
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
	models.DeleteEmployee(uid, companyID)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}


