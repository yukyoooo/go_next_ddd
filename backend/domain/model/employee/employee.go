package employee

import (
	model "github.com/yukyoooo/go_next_ddd/domain/model"
	"github.com/yukyoooo/go_next_ddd/enum"
)

type Employee struct {
	ID       int
	Name     FullName
	Email    Email
	Password Password
	Role     enum.Role
}

func NewEmployee(name FullName, email Email, password Password, role enum.Role) (*Employee, error) {
	return &Employee{Name: name, Email: email, Password: password, Role: role}, nil
}

func (e *Employee) Save() (err error) {
	employeeRepository := NewEmployeeRepository(model.Db)
	err = employeeRepository.Save(e)
	if err != nil {
		return err
	}
	return err
}

func GetEmployee(id int) (employee *Employee, err error) {
	employeeRepository := NewEmployeeRepository(model.Db)
	employee, err = employeeRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return employee, err
}

func (e *Employee) UpdateUser() (err error) {
	employeeRepository := NewEmployeeRepository(model.Db)
	err = employeeRepository.Update(e)
	if err != nil {
		return err
	}
	return err
}

func (e *Employee) DeleteEmployee() (err error) {
	employeeRepository := NewEmployeeRepository(model.Db)
	err = employeeRepository.Remove(e.ID)
	if err != nil {
		return err
	}
	return err
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
