package employee

type EmployeeService struct {
	employeeRepository EmployeeRepository
}

func NewEmployeeService(employeeRepository EmployeeRepository) (*EmployeeService, error) {
	return &EmployeeService{employeeRepository: employeeRepository}, nil
}

func (es *EmployeeService) Exists(employee *Employee) (isExists bool, err error) {
	employees, err := es.employeeRepository.FindByNameAndEmail(employee.Name.FirstName, employee.Name.LastName, employee.Email.Value)
	if err != nil {
		return false, err
	}
	return employees != nil, nil
}
