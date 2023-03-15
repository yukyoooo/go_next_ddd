package employee

import (
	"database/sql"
	"fmt"
)

type EmployeeService struct {
	Db *sql.DB
}

func NewEmployeeService(Db *sql.DB) (*EmployeeService, error) {
	return &EmployeeService{Db: Db}, nil
}

func (employeeService *EmployeeService) Exists(employee *Employee) (isExists bool, err error) {
	tx, err := employeeService.Db.Begin()
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

	rows, err := tx.Query(cmd, employee.Name.firstName, employee.Name.lastName, employee.Email.value)
	if err != nil {
		return false, fmt.Errorf("userservice.Exists(): %v", err)
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, nil
}
