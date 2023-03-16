package employee

import (
	"log"

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

func (e *Employee) Save() (err error) {
	employeeRepository, err := NewEmployeeRepository(model.Db)
	if err != nil {
		return err
	}
	err = employeeRepository.Save(e)
	if err != nil {
		return err
	}
	return err
}

func GetEmployee(id int) (employee *Employee, err error) {
	employeeRepository, err := NewEmployeeRepository(model.Db)
	if err != nil {
		return nil, err
	}
	employee, err = employeeRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return employee, err
}

func (e *Employee) UpdateUser() (err error) {
	cmd := `update employees set first_name = ?, last_name = ?, email = ?, password = ?, role = ? where id = ?`
	_, err = model.Db.Exec(cmd, e.Name.firstName, e.Name.lastName, e.Email.value, e.Password.value, e.Role, e.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (e *Employee) DeleteEmployee() (err error) {
	cmd := `delete from employees where id = ?`
	_, err = model.Db.Exec(cmd, e.ID)
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
