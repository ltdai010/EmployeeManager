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
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"strconv"
	"strings"
)

var client StringBigsetService.Client
// server
type HandleCompany struct {

}

//get an employee
func (this *HandleCompany) GetEmployee(ctx context.Context, id string, companyID string) (*company.Employee, error) {
	_, e, err := checkExist(companyID, id)
	if err != nil {
		return nil, err
	}
	return e, nil
}

//post an employee
func (this *HandleCompany) PostEmployee(ctx context.Context, id string, name string, address string,
	date *company.Date, companyID string) (err error)  {
	if !validateDate(date) {
		return errors.New("the birthday is invalid")
	}
	mm := strconv.Itoa(int(date.GetMonth()))
	yyyy := strconv.Itoa(int(date.GetYear()))
	key := mm + "/" + yyyy + "-" + companyID
	_, i, err := checkExist(companyID, id)
	if i != nil {
		return errors.New("all ready exist")
	}
	c, err := client.BsGetItem2(models.CompanyKey, generic.TItemKey(companyID))
	if c == nil {
		return errors.New("company not exist")
	}
	s, err := json.Marshal(company.Employee{ID: id, Name: name,
		Address: address, Date: date, CompanyID: companyID})
	item := generic.TItem{Key: []byte(id), Value: s}
	_, err = client.BsPutItem2(generic.TStringKey(key), &item)
	return err
}

//put an employee
func (this *HandleCompany) PutEmployee(ctx context.Context, id string, name string, address string,
	date *company.Date, companyID string) (err error)  {
	k, e, err := checkExist(companyID, id)
	if e == nil {
		return errors.New("not exist")
	}
	_, err = client.BsRemoveItem2(k, generic.TItemKey(id))
	key := string(date.GetMonth()) + "/" + string(date.GetYear()) + "-" + companyID
	s, err := json.Marshal(company.Employee{ID: id, Name: name,
		Address: address, Date: date, CompanyID: companyID})
	item := generic.TItem{Key: []byte(id), Value: s}
	_, err = client.BsPutItem2(generic.TStringKey(key), &item)
	if err != nil {
		return err
	}
	return nil
}

//remove employee
func (this *HandleCompany) RemoveEmployee(ctx context.Context, id string, companyID string) (err error) {
	k, e, err := checkExist(companyID, id)
	if e == nil {
		return errors.New("not exist")
	}
	_, err = client.BsRemoveItem2(k, generic.TItemKey(id))
	return err
}

//get a company
func (this *HandleCompany) GetCompany(ctx context.Context, id string)  (*company.Company, error) {
	c, err := client.BsGetItem2(models.CompanyKey, generic.TItemKey(id))
	if c == nil {
		return nil, errors.New("not found")
	}
	var com company.Company
	err = json.Unmarshal(c.GetValue(), &com)
	if err != nil {
		return nil, err
	}
	return &com, err
}

//post a company
func (this *HandleCompany) PostCompany(ctx context.Context, id string, name string, address string) (err error) {
	i, err := client.BsGetItem2(models.CompanyKey, generic.TItemKey(id))
	if i != nil {
		return errors.New("already exist")
	}
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
func (this *HandleCompany) GetEmployeeList(ctx context.Context, companyID string) ([]*company.Employee, error) {
	var list []*company.Employee
	length, err := client.TotalStringKeyCount2()
	listk, err := client.GetListKey2(0, int32(length))
	for _, i := range listk{
		set := strings.Split(i, "-")
		if len(set) != 2 {
			continue
		}
		if set[1] == companyID {
			length, err = client.GetTotalCount2(generic.TStringKey(i))
			if err != nil {
				return nil, err
			}
			listi, err := client.BsGetSlice2(generic.TStringKey(i), 0, int32(length))
			if err != nil {
				return nil, err
			}
			for _, item := range listi{
				var em company.Employee
				err := json.Unmarshal(item.GetValue(), &em)
				if err != nil {
					return nil, err
				}
				list = append(list, &em)
			}
		}
	}
	if list == nil {
		return nil, errors.New("company not found")
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


//get all employee
func (this *HandleCompany) GetAllEmployee(ctx context.Context) (list []*company.Employee ,err error) {
	length, err := client.GetTotalCount2(models.CompanyKey)
	listcom, err := client.BsGetSlice2(models.CompanyKey, 0, int32(length))
	for _, i := range listcom{
		var c company.Company
		err = json.Unmarshal(i.GetValue(), &c)
		length, err = client.TotalStringKeyCount2()
		if err != nil {
			return nil, err
		}
		listk, err := client.GetListKey2(0, int32(length))
		if err != nil {
			return nil, err
		}
		for _, j := range listk {
			set := strings.Split(j, "-")
			if len(set) != 2 {
				continue
			}
			if set[1] == c.GetID() {
				length, err = client.GetTotalCount2(generic.TStringKey(j))
				if err != nil {
					return nil, err
				}
				listi, err := client.BsGetSlice2(generic.TStringKey(j), 0, int32(length))
				if err != nil {
					return nil, err
				}
				for _, item := range listi {
					var em company.Employee
					err := json.Unmarshal(item.GetValue(), &em)
					if err != nil {
						continue
					}
					list = append(list, &em)
				}
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return list, nil
}

//get list employee by date range
func (this *HandleCompany) GetListEmployeeInDate(ctx context.Context, companyID string, first *company.Date,
	last *company.Date) ([]*company.Employee, error) {
	if !validateDate(first) || !validateDate(last) {
		return nil, errors.New("date is invalid")
	}
	var list []*company.Employee
	length, err := client.TotalStringKeyCount2()
	listk, err := client.GetListKey2(0, int32(length))
	for _, i := range listk {
		set := strings.Split(i, "-")
		if len(set) != 2 {
			continue
		}
		if set[1] == companyID {
			setD := strings.Split(set[0], "/")
			log.Println(setD)
			if len(setD) != 2 {
				continue
			}
			log.Println(setD)
			mm, err := strconv.Atoi(setD[0])
			if err != nil {
				continue
			}
			yyyy, err := strconv.Atoi(setD[1])
			if err != nil {
				continue
			}
			date := company.Date{Day: 1, Month: company.Int(mm), Year: company.Int(yyyy)}
			if compareDate(&date, first) < 0 || compareDate(&date, last) > 0 {
				continue
			}
			length, err := client.GetTotalCount2(generic.TStringKey(i))
			if err != nil {
				continue
			}
			listi, err := client.BsGetSlice2(generic.TStringKey(i), 0, int32(length))
			if err != nil {
				continue
			}
			for _, item := range listi {
				var em company.Employee
				err := json.Unmarshal(item.GetValue(), &em)
				if err != nil {
					continue
				}
				if compareDate(em.GetDate(), first) >= 0 && compareDate(em.GetDate(), last) <= 0 {
					list = append(list, &em)
				}
			}
		}
 	}
	if list == nil {
		return nil, errors.New("no employee in this range")
	}
	if err != nil {
		return nil, err
	}
	return mergeSort(list), nil
}

func checkExist(companyID string, employeeID string) (generic.TStringKey ,*company.Employee, error)  {
	length, err := client.TotalStringKeyCount2()
	if err != nil {
		return "", nil, err
	}
	listk, err := client.GetListKey2(0, int32(length))
	for _, i := range listk {
		set := strings.Split(i, "-")
		if len(set) != 2 {
			continue
		}
		if set[1] == companyID {
			item, err := client.BsGetItem2(generic.TStringKey(i), generic.TItemKey(employeeID))
			if err != nil {
				continue
			}
			if item == nil {
				continue
			}
			var e company.Employee
			err = json.Unmarshal(item.GetValue(), &e)
			if err != nil {
				continue
			}
			return generic.TStringKey(i), &e, nil
		}
	}
	return "", nil, errors.New("not found")
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