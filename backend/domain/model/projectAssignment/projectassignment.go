package projectassignment

type ProjectAssignment struct {
	ID         int
	ProjectID  int
	EmployeeID int
}

func NewProjectAssignment(projectID int, employeeID int) (*ProjectAssignment, error) {
	return &ProjectAssignment{
		ProjectID:  projectID,
		EmployeeID: employeeID,
	}, nil
}
