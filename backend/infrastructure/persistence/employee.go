// package persistence

// import (
// 	"database/sql"

// 	model "github.com/yukyoooo/go_next_ddd/domain/model"
// 	"github.com/yukyoooo/go_next_ddd/domain/model/employee"
// 	"github.com/yukyoooo/go_next_ddd/domain/repository"
// )

// type employeePersistence struct {
// 	Db *sql.DB
// }

// func NewEmployeePersistence(Db *sql.DB) (*repository.EmployeeRepository, error) {
// 	return &employeePersistence{Db: Db}, nil
// }

// func (ep *employeePersistence) save(first_name string, last_name string, email string, password string, role int) (err error){
// 	cmd := `insert into employees (
// 		first_name,
// 		last_name,
// 		email,
// 		password,
// 		role) values (?, ?, ?, ?, ?)`
// 	_, err = model.Db.Exec(cmd, first_name, last_name, email, password, role)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (ep *employeePersistence) getById(id int) (*employee.Employee, error) {
// 	var e = employee.Employee{}
// 	cmd := `select id, first_name, last_name, email, password, role
// 	from employees where id = ?`
// 	err := model.Db.QueryRow(cmd, id).Scan(
// 		&e.ID,
// 		&e.Name,
// 		&e.Name,
// 		&e.Email,
// 		&e.Password,
// 		&e.Role,
// 	)

// 	return e, err
// }
// func (ep *employeePersistence) update(id int, first_name string, last_name string, email string, password string, role int)
// func (ep *employeePersistence) delete(id int)