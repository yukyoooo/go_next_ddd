package application

import (
	"fmt"
	"log"

	"github.com/yukyoooo/go_next_ddd/domain/model/employee"
	"github.com/yukyoooo/go_next_ddd/enum"
)

type EmployeeApplicationService struct {
	employeeRepository employee.EmployeeRepository
	employeeService    employee.EmployeeService
}

func NewEmployeeApplicationService(employeeRepository employee.EmployeeRepository, employeeService employee.EmployeeService) *EmployeeApplicationService {
	return &EmployeeApplicationService{
		employeeRepository: employeeRepository,
		employeeService:    employeeService,
	}
}

func (eas *EmployeeApplicationService) Register(firstName string, lastName string, email string, password string, role int) error {
	newEmployeeName, err := employee.NewFullName(firstName, lastName)
	if err != nil {
		return err
	}

	newEmail, err := employee.NewEmail(email)
	if err != nil {
		return err
	}

	newPassword, err := employee.NewPassword(password)
	if err != nil {
		return err
	}

	newEmployee, err := employee.NewEmployee(*newEmployeeName, *newEmail, *newPassword, enum.Role(role))
	if err != nil {
		return err
	}

	isExists, err := eas.employeeService.Exists(newEmployee)
	if err != nil {
		return err
	}

	if isExists {
		return fmt.Errorf("userservice.Exists() :既に存在する名前、もしくはメールアドレスです")
	}

	if err := eas.employeeRepository.Save(newEmployee); err != nil {
		return err
	}

	log.Println("employee is successfully added in employees table. employee:", newEmployee)
	return nil
}
