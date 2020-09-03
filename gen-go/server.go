package main

import (
	"company-manager/gen-go/company"
	"company-manager/models"
	"context"
	"encoding/json"
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
func (this *HandleCompany) GetEmployee(ctx context.Context, id string, companyID string) (r string, err error) {
	e, err := client.BsGetItem2(generic.TStringKey(companyID), []byte(id))
	if e == nil {
		return "failed", err
	}
	var s company.Employee
	err = json.Unmarshal(e.GetValue(), &s)
	i, _ := client.BsGetItem2(models.CompanyKey, []byte(s.GetCompany()))
	var c company.Company
	err = json.Unmarshal(i.GetValue(), &c)
	s.Company = c.GetName()
	if err != nil {
		return "failed", nil
	}
	return s.String(), nil
}

//post an employee
func (this *HandleCompany) PostEmployee(ctx context.Context, id string, name string, address string,
	age company.Int, companyID string) (err error)  {
	s, err := json.Marshal(company.Employee{ID: id, Name: name,
		Address: address, Age: age, Company: companyID})
	item := generic.TItem{Key: []byte(id), Value: s}
	_, err = client.BsPutItem2(generic.TStringKey(companyID), &item)
	return err
}

//put an employee
func (this *HandleCompany) PutEmployee(ctx context.Context, id string, name string, address string,
	age company.Int, companyID string) (err error)  {
	i, err := client.BsGetItem2(generic.TStringKey(companyID), generic.TItemKey(id))
	if i == nil {
		return err
	}
	s, err := json.Marshal(company.Employee{ID: id, Name: name,
		Address: address, Age: age, Company: companyID})
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
func (this *HandleCompany) GetCompany(ctx context.Context, id string)  (string string, err error) {
	c, err := client.BsGetItem2(models.CompanyKey, []byte(id))
	if c == nil {
		return "failed", err
	}
	var s company.Company
	err = json.Unmarshal(c.GetValue(), &s)
	if err != nil {
		return "failed", err
	}
	return s.String(), nil
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
func (this *HandleCompany) PutCompany(ctx context.Context, id string, name string, address string) (err error) {
	i, err := client.BsGetItem2(models.CompanyKey, generic.TItemKey(id))
	if i == nil {
		return err
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
func (this *HandleCompany) GetEmployeeList(ctx context.Context, id string) (list []string, err error) {
	count, err := client.GetTotalCount2(generic.TStringKey(id))
	arr, err := client.BsGetSlice2(generic.TStringKey(id), 0, int32(count))
	for _, i := range arr{
		var e company.Employee
		json.Unmarshal(i.GetValue(), &e)
		list = append(list, e.String())
	}
	return list, err
}

//get all company
func (this *HandleCompany) GetAllCompany(ctx context.Context) (list []string, err error){
	count, err := client.GetTotalCount2(models.CompanyKey)
	listItem, err := client.BsGetSlice2(models.CompanyKey, 0, int32(count))
	for _, i := range listItem {
		var c company.Company
		json.Unmarshal(i.GetValue(), &c)
		list = append(list, c.String())
	}
	return list, err
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