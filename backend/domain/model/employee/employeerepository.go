package employee

type EmployeeRepository interface {
	Save(employee *Employee) error
	FindById(id int) (*Employee, error)
	Update(employee *Employee) error
	Remove(id int) error
	FindByNameAndEmail(first_name string, last_name string, email string) (*Employee, error)
}
