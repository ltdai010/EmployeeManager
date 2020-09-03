package models

import (
	"company-manager/gen-go/company"
)

func init() {

}



func AddUser(u company.Employee) string {
	err := client.PutEmployee(defaultContext, u.ID, u.Name, u.Address, u.Age, u.Company)
	if err != nil {
		return err.Error()
	}
	return u.ID
}

func GetUser(uid string, companyID string) (u string, err error) {
	s, err := client.GetEmployee(defaultContext, uid, companyID)
	return s, err
}

func GetAllUsers(id string) (list []string) {
	list, err := client.GetEmployeeList(defaultContext, id)
	if err != nil {
		return nil
	}
	return list
}

func UpdateUser(uid string, name string, address string, age int, companyID string)  (err error) {
	err = client.PutEmployee(defaultContext, uid, name, address, company.Int(age), companyID)
	return err
}

func DeleteUser(uid string, companyID string) {
	client.RemoveEmployee(defaultContext, uid, companyID)
}
