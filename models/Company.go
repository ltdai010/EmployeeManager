package models

import (
	"company-manager/gen-go/company"
	"encoding/json"
	"errors"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)


func init() {

}

const CompanyKey = "Company"

func AddOne(company company.Company) string {
	v, _ := json.Marshal(company)
	item := generic.TItem{Key: []byte(company.ID), Value: v}
	_, _ = client.BsPutItem2(CompanyKey, &item)
	return company.ID
}

func GetOne(companyID string) (*company.Company, error) {
	item, err := client.BsGetItem2(CompanyKey, generic.TItemKey(companyID))
	if err != nil {
		return nil, errors.New("ObjectId Not Exist")
	}
	var c company.Company
	_ = json.Unmarshal(item.GetValue(), &c)
	return &c, nil
}

func GetAll() []*company.Company {
	var list []*company.Company
	length, _ := client.GetTotalCount2(CompanyKey)
	listKey, _ := client.BsGetSlice2(CompanyKey, 0, int32(length))
	for _, k := range listKey {
		var c company.Company
		_ = json.Unmarshal(k.GetValue(), &c)
		list = append(list, &c)
	}
	return list
}

func Update(id string, name string, address string, list []string) error {
	item, err := client.BsGetItem2(CompanyKey, generic.TItemKey(id))
	if err != nil{
		return errors.New("employee Not Exist")
	}
	var c company.Company
	_ = json.Unmarshal(item.GetValue(), &c)
	if name != ""{
		c.Name = name
	}
	if address != ""{
		c.Address = address
	}
	if list != nil{
		c.EmployeeList = list
	}
	s, _ := json.Marshal(c)
	it := generic.TItem{Key: generic.TItemKey(id), Value: s}
	_, _ = client.BsPutItem2(CompanyKey, &it)
	return  nil
}

func Delete(ObjectId string) {
	_, _ = client.BsRemoveItem2(CompanyKey, generic.TItemKey(ObjectId))
}

