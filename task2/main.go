package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const DatabaseUser = "root"
const DatabasePassword = "password"

//	1)----------------------------------------------------------------------------------------------------------------------------------
const ManagerQuery = `	select
							t.title,
							e.first_name,
							e.last_name,
							s.salary
						from (select * from dept_manager where to_date = (select MAX(to_date) from dept_manager)) d
						join (select * from titles where to_date = (select MAX(to_date) from titles)) t on d.emp_no = t.emp_no
						join (select * from salaries where to_date = (select MAX(to_date) from salaries)) s on d.emp_no = s.emp_no
						join employees e on d.emp_no = e.emp_no;`

type manager struct {
	title     string
	firstName string
	lastName  string
	salary    int
}

func getAllCurrentManagers(db *sql.DB) []manager {
	rows, err := db.Query(ManagerQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	managers := []manager{}

	for rows.Next() {
		m := manager{}
		err := rows.Scan(&m.title, &m.firstName, &m.lastName, &m.salary)
		if err != nil {
			panic(err)
		}
		managers = append(managers, m)
	}

	return managers
}

//	2)--------------------------------------------------------------------------------------------------------------------------------------------------
const EmployeesQuery = `select
							departments.dept_name,
							t.title,
							employees.first_name,
							employees.last_name,
							employees.hire_date,
							YEAR(CURRENT_DATE()) - YEAR(employees.hire_date) as work_years
						from employees
						join (select * from dept_emp where to_date = (select MAX(to_date) from dept_emp)) dept_emp on employees.emp_no = dept_emp.emp_no
						join departments on departments.dept_no = dept_emp.dept_no
						join (select * from titles where to_date = (select MAX(to_date) from titles)) t on employees.emp_no = t.emp_no
						where MONTH(employees.hire_date) = MONTH(CURRENT_DATE());`

type employee struct {
	department string
	title      string
	firstName  string
	lastName   string
	hireDate   string
	workYears  int
}

func getAnniversaryEmployees(db *sql.DB) []employee {
	rows, err := db.Query(EmployeesQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	employees := []employee{}

	for rows.Next() {
		e := employee{}
		err := rows.Scan(&e.department, &e.title, &e.firstName, &e.lastName, &e.hireDate, &e.workYears)
		if err != nil {
			panic(err)
		}
		employees = append(employees, e)
	}

	return employees
}

//	3)-------------------------------------------------------------------------------------------------------------------------------
const DepartmentsQuery = `	select
								departments.dept_name,
								COUNT(*) as employee_count,
								SUM(s.salary) as sum_salary
							from departments
							join (select emp_no, dept_no from dept_emp where to_date = (select MAX(to_date) from dept_emp)) dept_emp on departments.dept_no = dept_emp.dept_no
							join employees e on dept_emp.emp_no = e.emp_no
							join (select emp_no, salary from salaries where to_date = (select MAX(to_date) from salaries)) s on e.emp_no = s.emp_no
							group by departments.dept_name;`

type department struct {
	name          string
	employeeCount int
	sumSalary     int
}

func getDepartments(db *sql.DB) []department {
	rows, err := db.Query(DepartmentsQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	departments := []department{}

	for rows.Next() {
		d := department{}
		err := rows.Scan(&d.name, &d.employeeCount, &d.sumSalary)
		if err != nil {
			panic(err)
		}
		departments = append(departments, d)
	}

	return departments
}

func main() {
	db, err := sql.Open("mysql", DatabaseUser+":"+DatabasePassword+"@/employees")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	managers := getAllCurrentManagers(db)
	for _, m := range managers {
		fmt.Println(m)
	}

	employees := getAnniversaryEmployees(db)
	for _, e := range employees {
		fmt.Println(e)
	}

	departments := getDepartments(db)
	for _, d := range departments {
		fmt.Println(d)
	}

}
