package models

import (
	"company-manager/gen-go/company"
)

func init() {

}



func AddEmployee(u company.Employee) string {
	err := client.PostEmployee(defaultContext, u.GetID(), u.GetName(), u.GetAddress(), u.GetDate(), u.GetCompanyID())
	if err != nil {
		return err.Error()
	}
	return u.GetID()
}

func GetEmployee(uid string, companyID string) (u *company.Employee, err error) {
	s, err := client.GetEmployee(defaultContext, uid, companyID)
	return s, err
}

func GetAllCompanyEmployee(id string) (list []*company.Employee) {
	list, err := client.GetEmployeeList(defaultContext, id)
	if err != nil {
		return nil
	}
	return list
}

func GetEmployeeSlice(companyID string, start int, count int) (list []*company.Employee, err error) {
	list, err = client.GetListEmployee(defaultContext, companyID, company.Int(start), company.Int(count))
	return list, err
}

func GetEmployeeSliceInTime(companyID string, first *company.Date,
	last *company.Date) (list []*company.Employee, err error) {
	list, err = client.GetListEmployeeInDate(defaultContext, companyID, first, last)
	return list, err
}

func GetAllEmployee() (list []*company.Employee, err error) {
	list, err = client.GetAllEmployee(defaultContext);
	return list, err
}

func UpdateEmployee(uid string, name string, address string, date *company.Date, companyID string)  (err error) {
	err = client.PutEmployee(defaultContext, uid, name, address, date, companyID)
	return err
}

func DeleteEmployee(uid string, companyID string) error {
	err := client.RemoveEmployee(defaultContext, uid, companyID)
	return err
}
