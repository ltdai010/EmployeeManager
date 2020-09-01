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
	fmt.Println("2: Put a new Employee")
	fmt.Println("3: Delete an Employee")
	fmt.Println("4: Get a Company information")
	fmt.Println("5: Put a new Company")
	fmt.Println("6: Delete a Company")
	fmt.Println("7: Get a Company Employee list")
	fmt.Print("Enter command: ")
	var command int
	_, _ = fmt.Scanf("%d", &command)
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
			fmt.Print("Enter an Employee id name address age companyID:")
			var id, name, address, companyID string
			var age int
			_, err = fmt.Scanf("%s%s%s%d%s", &id, &name, &address, &age, &companyID)
			err := client.PutEmployee(defaultCtx, id, name, address, company.Int(age), companyID)
			if err != nil {
				log.Println("Something wrong")
			} else {
				fmt.Println("Success")
			}
		case 3:
			fmt.Print("Enter the Employee id you want to delete:")
			var id string
			_, err = fmt.Scanf("%s", &id)
			err = client.RemoveEmployee(defaultCtx, id)
		case 4:
			fmt.Print("Enter Company id: ")
			var id string
			_, err = fmt.Scanf("%s", &id)
			s, err := client.GetCompany(defaultCtx, id)
			if err != nil {
				log.Println("Something wrong")
			} else {
				fmt.Println(s)
			}
		case 5:
			fmt.Print("Enter a Company id name address :")
			var id, name, address string
			_, err = fmt.Scanf("%s%s%s", &id, &name, &address)
			fmt.Println("Enter number of employee :")
			var num int
			_, err = fmt.Scanf("%d", &num)
			var list []string
			for i:=0; i < num; i++ {
				var s string
				_, err = fmt.Scanf("%s", &s)
				list = append(list, s)
			}
			err := client.PutCompany(defaultCtx, id, name, address, list)
			if err != nil {
				log.Println("Something wrong")
			} else {
				fmt.Println("Success")
			}
		case 6:
			fmt.Print("Enter the Company id you want to delete:")
			var id string
			_, err = fmt.Scanf("%s", &id)
			err = client.RemoveCompany(defaultCtx, id)
		case 7:
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
