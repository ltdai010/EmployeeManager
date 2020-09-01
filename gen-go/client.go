package main

import (
	"company-manager/gen-go/company"
	"context"
	"fmt"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
)

var defaultCtx = context.Background()

func handleClient(client *company.CompanyManagerClient) (err error) {
	fmt.Println("Command list:")
	fmt.Println("1: Get an Employee information")
	fmt.Println("2: Get a Company information")
	fmt.Println("3: Get a Company Employee list")
	fmt.Print("Enter command: ")
	var command int
	fmt.Scanf("%d", &command)
	switch command {
		case 1:
			fmt.Print("Enter Employee id: ")
			var id string
			_, err = fmt.Scanf("%s", &id)
			s, err := client.GetEmployee(defaultCtx, id)
			if err != nil {
				log.Println("Something wrong")
			} else {
				fmt.Println(s)
			}
		case 2:
			fmt.Print("Enter Company id: ")
			var id string
			_, err = fmt.Scanf("%s", &id)
			s, err := client.GetCompany(defaultCtx, id)
			if err != nil {
				log.Println("Something wrong")
			} else {
				fmt.Println(s)
			}
		case 3:
			fmt.Print("Enter Company id: ")
			var id string
			_, err = fmt.Scanf("%s", &id)
			s, err := client.GetEmployeeList(defaultCtx, id)
			if err != nil {
				log.Println("Something wrong")
			} else {
				fmt.Println(s)
			}
		default:
			fmt.Println("Wrong command")
	}
	return err
}

func main() {
	transport, err := thrift.NewTSocket("127.0.0.1:8888")
	if err != nil {
		log.Println(err.Error())
		return
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	err = handleClient(company.NewCompanyManagerClient(thrift.NewTStandardClient(iprot, oprot)))
	if err != nil {
		fmt.Println("Error", err)
	}
}
