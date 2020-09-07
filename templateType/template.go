package templateType

import (
	"company-manager/gen-go/company"
)

type UpdateEmployeeForm struct {
	Name        string
	DateOfBirth *company.Date
	Address     string
}

type UpdateCompanyForm struct {
	Name string
	Address string
}
