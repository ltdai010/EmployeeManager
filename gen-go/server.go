package main

import (
	"company-manager/gen-go/company"
	"company-manager/models"
	"context"
	"encoding/json"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
)

var client = StringBigsetService.NewStringBigsetServiceModel("my-service-id",
	[]string{"0.0.0.0:2360"},
	GoEndpointBackendManager.EndPoint{
		Host:      "127.0.0.1",
		Port:      "18407",
		ServiceID: "my-service-id",
	})
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

func (this *HandleCompany) GetCompany(ctx context.Context, id string)  (string, error) {
	c, err := client.BsGetItem2(models.CompanyKey, []byte(id))
	if err != nil {
		return "", err
	}
	var s company.Company
	_ = json.Unmarshal(c.GetValue(), &s)
	return s.String(), nil
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
