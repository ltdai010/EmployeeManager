namespace go company

typedef i32 int

struct Employee{
    1: string id
    2: string name
    3: int age
    4: string address
    5: string company
}

struct Company{
    1: string id
    2: string name
    3: string address
    4: list<string> employeeList 
}

service companyManager{
    string getEmployee(1:string id)
    string getCompany(1:string id)
    list<string> getEmployeeList(1:string id)
}
