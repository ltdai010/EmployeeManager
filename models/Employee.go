package models

import (
	"company-manager/gen-go/company"
	"encoding/json"
	"errors"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

var (
	client StringBigsetService.Client
)

const EmployeeKey = "Employee"

func init() {
	client       = StringBigsetService.NewStringBigsetServiceModel("my-service-id",
		[]string{"0.0.0.0:2379"},
		GoEndpointBackendManager.EndPoint{
			Host:      "127.0.0.1",
			Port:      "18407",
			ServiceID: "my-service-id",
		})
	_, _ = client.CreateStringBigSet2(EmployeeKey)
	_, _ = client.CreateStringBigSet2(CompanyKey)
}



func AddUser(u company.Employee) string {
	v, _ := json.Marshal(u)
	item := generic.TItem{Key: []byte(u.ID), Value: v}
	_, _ = client.BsPutItem2(EmployeeKey, &item)
	return u.ID
}

func GetUser(uid string) (u *company.Employee, err error) {
	item, err := client.BsGetItem2(EmployeeKey, []byte(uid))
	if err != nil{
		return nil, errors.New("employee not exists")
	}
	var e company.Employee
	_ = json.Unmarshal(item.GetValue(), &e)
	return &e, nil
}

func GetAllUsers() []*company.Employee {
	var list []*company.Employee
	length, _ := client.GetTotalCount2(EmployeeKey)
	listKey, _ := client.BsGetSlice2(EmployeeKey, 0, int32(length))
	for _, k := range listKey {
		var e company.Employee
		_ = json.Unmarshal(k.GetValue(), &e)
		list = append(list, &e)
	}
	return list
}

func UpdateUser(uid string, name string, address string, age int, companyID string)  error {
	item, err := client.BsGetItem2(EmployeeKey, generic.TItemKey(uid))
	if err != nil{
		return errors.New("employee Not Exist")
	}
	var e company.Employee
	_ = json.Unmarshal(item.GetValue(), &e)
	if name != ""{
		e.Name = name
	}
	if address != ""{
		e.Address = address
	}
	if age != 0{
		e.Age = company.Int(age)
	}
	if companyID != ""{
		e.Company = companyID
	}
	s, _ := json.Marshal(e)
	it := generic.TItem{Key: generic.TItemKey(uid), Value: s}
	_, _ = client.BsPutItem2(EmployeeKey, &it)
	return nil
}


func DeleteUser(uid string) {
	client.BsRemoveItem2(EmployeeKey, generic.TItemKey(uid))
}
