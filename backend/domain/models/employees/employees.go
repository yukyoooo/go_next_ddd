package employee

import (
	"log"

	"github.com/yukyoooo/go_next_ddd/domain/models"
	"github.com/yukyoooo/go_next_ddd/enum"
)

type Employee struct {
	ID       int
	Name     FullName
	Email    Email
	Password Password
	Role     enum.Role
}

func (e *Employee) CreateEmployee() (err error) {
	cmd := `insert into employees (
		first_name,
		last_name,
		email,
		password,
		role) values (?, ?, ?, ?, ?)`

	_, err = models.Db.Exec(
		cmd,
		e.Name.firstName,
		e.Name.lastName,
		e.Email.value,
		e.Password.value,
		e.Role)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetEmployee(id int) (employee Employee, err error) {
	employee = Employee{}
	cmd := `select id, first_name, last_name, email, password, role
	from employees where id = ?`
	err = models.Db.QueryRow(cmd, id).Scan(
		&employee.ID,
		&employee.Name.firstName,
		&employee.Name.lastName,
		&employee.Email.value,
		&employee.Password.value,
		&employee.Role,
	)

	return employee, err
}

func (e *Employee) UpdateUser() (err error) {
	cmd := `update employees set first_name = ?, last_name = ?, email = ?, password = ?, role = ? where id = ?`
	_, err = models.Db.Exec(cmd, e.Name.firstName, e.Name.lastName, e.Email.value, e.Password.value, e.Role, e.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (e *Employee) DeleteEmployee() (err error) {
	cmd := `delete from employees where id = ?`
	_, err = models.Db.Exec(cmd, e.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func NewEmployee(name FullName, email Email, password Password, role enum.Role) (*Employee, error) {
	return &Employee{Name: name, Email: email, Password: password, Role: role}, nil
}

func (e *Employee) WithChangeFirstName(firstName string) (_ *FullName, err error) {
	changedFullName, err := NewFullName(firstName, e.Name.lastName)
	if err != nil {
		return nil, err
	}
	e.Name = *changedFullName
	return changedFullName, nil
}

func (e *Employee) WithChangeLastName(lastName string) (_ *FullName, err error) {
	changedFullName, err := NewFullName(e.Name.firstName, lastName)
	if err != nil {
		return nil, err
	}
	e.Name = *changedFullName
	return changedFullName, nil
}
