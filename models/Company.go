package models

import (
	"company-manager/gen-go/company"
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/astaxie/beego"
	"log"
)

var (
	defaultContext = context.Background()
	client *company.CompanyManagerClient
)

func init() {
	transport, err := thrift.NewTSocket(beego.AppConfig.String("hostPort"))
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
	log.Println("ok")
	err := client.PostCompany(defaultContext, company.GetID(), company.GetName(), company.GetAddress())
	if err != nil {
		return err.Error()
	}
	return company.ID
}

func GetOne(companyID string) (*company.Company, error) {
	s, err := client.GetCompany(defaultContext, companyID)
	log.Println(companyID)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func GetAll() []*company.Company {
	list, err := client.GetAllCompany(defaultContext)
	if err != nil {
		return nil
	}
	return list
}

func Update(id string, name string, address string) error {
	err := client.PutCompany(defaultContext, id, name, address)
	return err
}

func Delete(ObjectId string) error{
	err := client.RemoveCompany(defaultContext, ObjectId)
	if err != nil {
		return err
	}
	return nil
}

