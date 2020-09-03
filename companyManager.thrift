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
}


service companyManager{
    string getEmployee(1:string id, 2:string companyID)
    void postEmployee(1:string id, 2:string name, 3:string address, 4: int age, 5: string company)
    void putEmployee(1:string id, 2:string name, 3:string address, 4: int age, 5: string company)
    void removeEmployee(1:string id, 2:string companyID)
    string getCompany(1:string id)
    list<string> getAllCompany()
    void postCompany(1:string id, 2:string name, 3:string address)
    void putCompany(1:string id, 2:string name, 3:string address)
    list<string> getEmployeeList(1:string id)
    void removeCompany(1:string id)
}
