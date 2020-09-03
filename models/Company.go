package models

import (
	"company-manager/gen-go/company"
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
)

var (
	defaultContext = context.Background()
	client *company.CompanyManagerClient
)

func init() {
	transport, err := thrift.NewTSocket("127.0.0.1:8888")
	if err != nil {
		log.Println(err.Error())
		return
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	if err := transport.Open(); err != nil {
		return
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client = company.NewCompanyManagerClient(thrift.NewTStandardClient(iprot, oprot))
}


const CompanyKey = "Company"

func AddOne(company company.Company) string {
	err := client.PutCompany(defaultContext, company.ID ,company.Name, company.Address)
	if err != nil {
		return err.Error()
	}
	return company.ID
}

func GetOne(companyID string) (string, error) {
	s, err := client.GetCompany(defaultContext, companyID)
	return s, err
}

func GetAll() []string {
	list, err := client.GetAllCompany(defaultContext)
	if err != nil {
		return nil
	}
	return list
}

func Update(id string, name string, address string) error {
	return  client.PutCompany(defaultContext, id, name, address)
}

func Delete(ObjectId string) error{
	err := client.RemoveCompany(defaultContext, ObjectId)
	if err != nil {
		return err
	}
	return nil
}

