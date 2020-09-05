package templateType

import (
	"company-manager/gen-go/company"
)

type updateEmployeeForm struct {
	name string
	dateOfBirth company.Date
	address string
}

type updateCompanyForm struct {
	name string
	address string
}
