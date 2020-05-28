
-- 1) Find all current managers of each department and display his/her title, first name, last name, current salary.
select
       t.title,
       e.first_name,
       e.last_name,
       s.salary
from (select * from dept_manager where to_date = (select MAX(to_date) from dept_manager)) d
join (select * from titles where to_date = (select MAX(to_date) from titles)) t on d.emp_no = t.emp_no
join (select * from salaries where to_date = (select MAX(to_date) from salaries)) s on d.emp_no = s.emp_no
join employees e on d.emp_no = e.emp_no;


-- 2) Find all employees (department, title, first name, last name, hire date, how many years they have been working) to congratulate them on their hire anniversary this month.
select
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
where MONTH(employees.hire_date) = MONTH(CURRENT_DATE());


-- 3) Find all departments, their current employee count, their current sum salary.
select
       departments.dept_name,
       COUNT(*) as employee_count,
       SUM(s.salary) as sum_salary
from departments
join (select emp_no, dept_no from dept_emp where to_date = (select MAX(to_date) from dept_emp)) dept_emp on departments.dept_no = dept_emp.dept_no
join employees e on dept_emp.emp_no = e.emp_no
join (select emp_no, salary from salaries where to_date = (select MAX(to_date) from salaries)) s on e.emp_no = s.emp_no
group by departments.dept_name;
