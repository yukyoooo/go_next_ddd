package employee

type EmployeeService struct {
	employeeRepository EmployeeRepositorier
}

func NewEmployeeService(employeeRepository EmployeeRepositorier) (*EmployeeService, error) {
	return &EmployeeService{employeeRepository: employeeRepository}, nil
}

func (es *EmployeeService) Exists(employee *Employee) (isExists bool, err error) {
	employees, err := es.employeeRepository.FindByNameAndEmail(employee.Name.firstName, employee.Name.lastName, employee.Email.value)
	if err != nil {
		return false, err
	}
	return employees != nil, nil
}