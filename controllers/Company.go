package controllers

import (
	"company-manager/gen-go/company"
	"company-manager/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about object
type CompanyController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body	body	company.Company	true	"The object content"
// @Success 200 {string} company.Company.ID
// @Failure 403 body is empty
// @router / [post]
func (o *CompanyController) Post() {
	var ob company.Company
	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	companyID := models.AddOne(ob)
	o.Data["json"] = map[string]string{"CompanyID": companyID}
	o.ServeJSON()
}

// @Title Get
// @Description find object by companyID
// @Param	companyID		query 	string	true		"the companyID you want to get"
// @Success 200 {object} company.Company
// @Failure 403 :objectId is empty
// @router /getCompany [get]
func (o *CompanyController) Get() {
	companyID := o.GetString("companyID")
	if companyID != "" {
		ob, err := models.GetOne(companyID)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} company.Company
// @Failure 403 :objectId is empty
// @router / [get]
func (o *CompanyController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	uid			query	string			true		"ID"
// @Param	body		body 	company.Company	true		"The body"
// @Success 200 {object} company.Company
// @Failure 403 :companyID is empty
// @router /update [put]
func (o *CompanyController) Put() {
	var ob company.Company
	id := o.GetString("uid")
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if id != "" {
		err = models.Update(id, ob.Name, ob.Address)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = "update success!"
		}
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	companyIO		query 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 companyID is empty
// @router /deleteCompany [delete]
func (o *CompanyController) Delete() {
	objectId := o.GetString("companyIO")
	err := models.Delete(objectId)
	if err != nil {
		o.Data["json"] = err
	} else {
		o.Data["json"] = "delete success!"
	}
	o.ServeJSON()
}

