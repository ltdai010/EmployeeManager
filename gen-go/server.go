package main

import (
	"company-manager/gen-go/company"
	"company-manager/models"
	"context"
	"encoding/json"
	"errors"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
)

var client StringBigsetService.Client
// server
type HandleCompany struct {

}

//get an employee
func (this *HandleCompany) GetEmployee(ctx context.Context, id string, companyID string) (e *company.Employee, err error) {
	item, err := client.BsGetItem2(generic.TStringKey(companyID), []byte(id))
	if item == nil {
		return nil, err
	}
	err = json.Unmarshal(item.GetValue(), e)
	if err != nil {
		return	nil, err
	}
	return e, nil
}

//post an employee
func (this *HandleCompany) PostEmployee(ctx context.Context, id string, name string, address string,
	date *company.Date, companyID string) (err error)  {
	if !validateDate(date) {
		return errors.New("the birthday is invalid")
	}
	s, err := json.Marshal(company.Employee{ID: id, Name: name,
		Address: address, Date: date, CompanyID: companyID})
	item := generic.TItem{Key: []byte(id), Value: s}
	_, err = client.BsPutItem2(generic.TStringKey(companyID), &item)
	return err
}

//put an employee
func (this *HandleCompany) PutEmployee(ctx context.Context, id string, name string, address string,
	date *company.Date, companyID string) (err error)  {
	i, err := client.BsGetItem2(generic.TStringKey(companyID), generic.TItemKey(id))
	if i == nil || err != nil{
		return errors.New("employee not exist")
	}
	s, err := json.Marshal(company.Employee{ID: id, Name: name,
		Address: address, Date: date, CompanyID: companyID})
	item := generic.TItem{Key: []byte(id), Value: s}
	_, err = client.BsPutItem2(generic.TStringKey(companyID), &item)
	return err
}

//remove employee
func (this *HandleCompany) RemoveEmployee(ctx context.Context, id string, companyID string) (err error) {
	_, err = client.BsRemoveItem2(generic.TStringKey(companyID), generic.TItemKey(id))
	return err
}

//get a company
func (this *HandleCompany) GetCompany(ctx context.Context, id string)  (com *company.Company, err  error) {
	c, err := client.BsGetItem2(models.CompanyKey, []byte(id))
	if c == nil {
		return nil, err
	}
	err = json.Unmarshal(c.GetValue(), c)
	if err != nil {
		return nil, err
	}
	return com, nil
}

//post a company
func (this *HandleCompany) PostCompany(ctx context.Context, id string, name string, address string) (err error) {
	c := company.Company{ID: id, Address: address, Name: name}
	s, err := json.Marshal(c)
	item := generic.TItem{Key: []byte(id), Value: s}
	_, err = client.BsPutItem2(models.CompanyKey, &item)
	return err
}

//put a company
func (this *HandleCompany) PutCompany(ctx context.Context, id string, name string, address string) error {
	i, err := client.BsGetItem2(models.CompanyKey, generic.TItemKey(id))
	if i == nil {
		return errors.New("company not exist")
	}
	c := company.Company{ID: id, Address: address, Name: name}
	s, err := json.Marshal(c)
	item := generic.TItem{Key: []byte(id), Value: s}
	_, err = client.BsPutItem2(models.CompanyKey, &item)
	return err
}

//remove a company
func (this *HandleCompany) RemoveCompany(ctx context.Context, id string) (err error) {
	_, err = client.BsRemoveItem2(models.CompanyKey, generic.TItemKey(id))
	return err
}


//get a company employeelist
func (this *HandleCompany) GetEmployeeList(ctx context.Context, id string) (list []*company.Employee, err error) {
	count, err := client.GetTotalCount2(generic.TStringKey(id))
	arr, err := client.BsGetSlice2(generic.TStringKey(id), 0, int32(count))
	for _, i := range arr {
		var e company.Employee
		err = json.Unmarshal(i.GetValue(), &e)
		list = append(list, &e)
	}
	if err != nil {
		return nil, err
	}
	return list, err
}

//get all company
func (this *HandleCompany) GetAllCompany(ctx context.Context) (list []*company.Company, err error){
	count, err := client.GetTotalCount2(models.CompanyKey)
	listItem, err := client.BsGetSlice2(models.CompanyKey, 0, int32(count))
	for _, i := range listItem {
		var c company.Company
		err = json.Unmarshal(i.GetValue(), &c)
		list = append(list, &c)
	}
	return list, err
}

//get list employee by range
func (this *HandleCompany) GetListEmployee(ctx context.Context, companyID string, start company.Int,
	count company.Int) (list []*company.Employee, err error) {
	listItem, err := client.BsGetSlice2(generic.TStringKey(companyID), int32(start), int32(count))
	for _, i := range listItem {
		var e company.Employee
		err = json.Unmarshal(i.GetValue(), &e)
		list = append(list, &e)
	}
	if err != nil {
		return nil, err
	}
	return list, nil
}

//get all employee
func (this *HandleCompany) GetAllEmployee(ctx context.Context) (list []*company.Employee ,err error) {
	length, err := client.GetTotalCount2(models.CompanyKey)
	listcom, err := client.BsGetSlice2(models.CompanyKey, 0, int32(length))
	for _, i := range listcom{
		var c company.Company
		err = json.Unmarshal(i.GetValue(), &c)
		length, err = client.GetTotalCount2(generic.TStringKey(c.GetID()))
		if err != nil {
			return nil, err
		}
		listemp, err := client.BsGetSlice2(generic.TStringKey(c.GetID()), 0, int32(length))
		if err != nil {
			return nil, err
		}
		for _, j := range listemp {
			var e company.Employee
			err = json.Unmarshal(j.GetValue(), &e)
			list = append(list, &e)
		}
	}
	if err != nil {
		return nil, err
	}
	return list, err
}

//get list employee by date range
func (this *HandleCompany) GetListEmployeeInDate(ctx context.Context, companyID string, first *company.Date,
	last *company.Date) (r []*company.Employee, err error) {
	if !validateDate(first) || !validateDate(last) {
		return nil, err
	}
	length, err := client.GetTotalCount2(generic.TStringKey(companyID))
	list, err := client.BsGetSlice2(generic.TStringKey(companyID), 0, int32(length))
	var unSortedList []*company.Employee
	for _, i := range list{
		var e company.Employee
		err = json.Unmarshal(i.GetValue(), &e)
		if compareDate(e.GetDate(), first) > 0 && compareDate(e.GetDate(), last) < 0 {
			unSortedList = append(unSortedList, &e)
		}
	}
	if err != nil {
		return nil, err
	}
	r = mergeSort(unSortedList)
	return r, nil
}

func validateDate(date *company.Date) bool {
	if date.GetYear() < 0 {
		return false
	} 
	if date.GetMonth() < 1 || date.GetMonth() > 12 {
		return false
	}
	if date.GetDay() < 1 {
		return false
	}
	switch date.Month {
	case 4, 6, 9, 11:
		if date.GetDay() > 30 {
			return false
		}
	case 2:
		if date.GetDay() > 29 {
			return false
		}
	default:
		if date.GetDay() > 31 {
			return false
		}
	}
	return true
}

func mergeSort(items []*company.Employee) []*company.Employee {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left = make([]*company.Employee, middle)
		right = make([]*company.Employee, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []*company.Employee) (result []*company.Employee) {
	result = make([]*company.Employee, len(left) + len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if compareDate(left[0].GetDate(), right[0].GetDate()) <= 0 {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

func compareDate(pro *company.Date, con *company.Date) int {
	result := compare(int(pro.GetYear()), int(con.GetYear())) 
	if result == 0{
		result = compare(int(pro.GetMonth()), int(con.GetMonth()))
		if result == 0 {
			result = compare(int(pro.Day), int(con.Day))
		}
	}
	return result
}

func compare(a int, b int) int  {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func main() {
	client = StringBigsetService.NewStringBigsetServiceModel("my-service-id",
		[]string{"0.0.0.0:2379"},
		GoEndpointBackendManager.EndPoint{
			Host:      "127.0.0.1",
			Port:      "18407",
			ServiceID: "my-service-id",
		})
	handle := &HandleCompany{}
	transport, err := thrift.NewTServerSocket("127.0.0.1:8888")
	if err != nil {
		log.Println(err.Error())
		return
	}
	processor := company.NewCompanyManagerProcessor(handle)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	err = server.Serve()
	if err != nil {
		log.Println(err.Error())
		return
	}
}