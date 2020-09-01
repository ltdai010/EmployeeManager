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

func (this *HandleCompany) GetEmployee(ctx context.Context, id string) (r string, err error) {
	e, err := client.BsGetItem2(models.EmployeeKey, []byte(id))
	if err != nil {
		return "", err
	}
	var s company.Employee
	_ = json.Unmarshal(e.GetValue(), &s)
	i, _ := client.BsGetItem2(models.CompanyKey, []byte(s.GetCompany()))
	var c company.Company
	_ = json.Unmarshal(i.GetValue(), &c)
	s.Company = c.GetName()
	return s.String(), nil
}

func (this *HandleCompany) PutEmployee(ctx context.Context, id string, name string, address string,
	age company.Int, companyID string) (err error)  {
	s, err := json.Marshal(company.Employee{ID: id, Name: name,
		Address: address, Age: age, Company: companyID})
	item := generic.TItem{Key: []byte(id), Value: s}
	_, err = client.BsPutItem2(models.EmployeeKey, &item)
	return err
}

func (this *HandleCompany) RemoveEmployee(ctx context.Context, id string) (err error) {
	_, err = client.BsRemoveItem2(models.EmployeeKey, generic.TItemKey(id))
	return err
}

func (this *HandleCompany) GetCompany(ctx context.Context, id string)  (string, error) {
	c, err := client.BsGetItem2(models.CompanyKey, []byte(id))
	if err != nil {
		return "", err
	}
	var s company.Company
	_ = json.Unmarshal(c.GetValue(), &s)
	return s.String(), nil
}

func (this *HandleCompany) PutCompany(ctx context.Context, id string, name string, address string,
	emplist []string) (err error) {
	c := company.Company{ID: id, Address: address, Name: name, EmployeeList: emplist}
	s, err := json.Marshal(c)
	item := generic.TItem{Key: []byte(id), Value: s}
	_, err = client.BsPutItem2(models.CompanyKey, &item)
	return err
}

func (this *HandleCompany) RemoveCompany(ctx context.Context, id string) (err error) {
	_, err = client.BsRemoveItem2(models.CompanyKey, generic.TItemKey(id))
	return err
}

func (this *HandleCompany) GetEmployeeList(ctx context.Context, id string) ([]string, error) {
	c, err := client.BsGetItem2(models.CompanyKey, []byte(id))
	if err != nil {
		return nil, err
	}
	var s company.Company
	_ = json.Unmarshal(c.GetValue(), &s)
	var list []string
	for _, i := range s.GetEmployeeList(){
		v, _ := client.BsGetItem2(models.EmployeeKey, []byte(i))
		var e company.Employee
		_ = json.Unmarshal(v.GetValue(), &e)
		list = append(list, e.Name)
	}
	return list, nil
}


func main() {
	client = StringBigsetService.NewStringBigsetServiceModel("my-service-id",
		[]string{"0.0.0.0:2360"},
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
