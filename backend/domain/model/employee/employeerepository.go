package employee

import (
	"database/sql"
	"fmt"

	model "github.com/yukyoooo/go_next_ddd/domain/model"
)

type EmployeeRepositorier interface {
	Save(employee *Employee) (error)
	FindById(id int) (*Employee, error)
	Update(employee *Employee) (error)
	Remove(id int) (error)
	FindByNameAndEmail(first_name string, last_name string, email string) (*Employee, error)
}

type EmployeeRepository struct {
	db *sql.DB
}



func NewEmployeeRepository(db *sql.DB) (*EmployeeRepository) {
	return &EmployeeRepository{db: db}
}

func (er *EmployeeRepository) Save(employee *Employee) (err error){
	cmd := `insert into employees (
		first_name,
		last_name,
		email,
		password,
		role) values (?, ?, ?, ?, ?)`
	_, err = model.Db.Exec(cmd, employee.Name.firstName, employee.Name.lastName, employee.Email.value, employee.Password.value, employee.Role)
	if err != nil {
		return err
	}
	return nil
}

func (er *EmployeeRepository) FindById(id int) (employee *Employee, err error) {
	cmd := `select id, first_name, last_name, email, password, role
	from employees where id = ?`
	employee = new(Employee)
	err = model.Db.QueryRow(cmd, id).Scan(
	&employee.ID,
		&employee.Name.firstName,
		&employee.Name.lastName,
		&employee.Email.value,
		&employee.Password.value,
		&employee.Role)
	if err != nil {
		return employee, err
	}
	
	return employee, err
}

func (er *EmployeeRepository) Update(employee *Employee) (err error) {
	cmd := `update employees set first_name = ?, last_name = ?, email = ?, password = ?, role = ? where id = ?`
	_, err = model.Db.Exec(cmd, employee.Name.firstName, employee.Name.lastName, employee.Email.value, employee.Password.value, employee.Role, employee.ID)
	if err != nil {
		return err
	}
	return nil
}

func (er *EmployeeRepository) Remove(id int) (err error) {
	cmd := `delete from employees where id = ?`
	_, err = model.Db.Exec(cmd, id)
	if err != nil {
		return err
	}
	return nil
}

func (er *EmployeeRepository) FindByNameAndEmail(firstName string, lastName string, email string) (employee *Employee, err error) {
	tx, err := er.db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	cmd := `select * from employees where (first_name = ? and last_name = ?) or email = ?`

	rows, err := tx.Query(cmd, firstName, lastName, email)
	if err != nil {
		return nil, fmt.Errorf("userservice.Exists(): %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&employee.ID,
			&employee.Name.firstName,
			&employee.Name.lastName,
			&employee.Email.value,
			&employee.Password.value,
			&employee.Role)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return employee, nil
}
