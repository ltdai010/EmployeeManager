namespace go company

typedef i32 int

struct Employee{
    1: string id
    2: string name
    3: string address
    4: string companyID
    5: Date date
}

struct Company{
    1: string id
    2: string name
    3: string address
}

struct Date{
    1: int day
    2: int month
    3: int year
}



service companyManager{
    Employee getEmployee(1:string id, 2:string companyID)
    void postEmployee(1:string id, 2:string name, 3:string address, 4:Date date, 5:string company)
    void putEmployee(1:string id, 2:string name, 3:string address, 4:Date date, 5:string company)
    void removeEmployee(1:string id, 2:string companyID)
    Company getCompany(1:string id)
    list<Employee> getAllEmployee()
    list<Employee> getListEmployeeInDate(1:string companyID, 2:Date first, 3:Date last)
    list<Company> getAllCompany()
    void postCompany(1:string id, 2:string name, 3:string address)
    void putCompany(1:string id, 2:string name, 3:string address)
    list<Employee> getEmployeeList(1:string id)
    void removeCompany(1:string id)
}
