package employee

import (
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
