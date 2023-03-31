package taskassignment

type TaskAssignment struct {
	ID         int
	TaskID     int
	EmployeeID int
}

func NewTaskAssignment(taskID int, employeeID int) (*TaskAssignment, error) {
	return &TaskAssignment{
		TaskID:     taskID,
		EmployeeID: employeeID,
	}, nil
}
