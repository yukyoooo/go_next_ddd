package projectassignment

import (
	"log"

	model "github.com/yukyoooo/go_next_ddd/domain/model"
)

type ProjectAssignment struct {
	ID          int
	ProjectID   int
	EmployeeID  int
}

func NewProjectAssignment(projectID int, employeeID int) (*ProjectAssignment, error) {
	return &ProjectAssignment{
		ProjectID:  projectID,
		EmployeeID: employeeID,
	}, nil
}

func (p *ProjectAssignment) Create() (err error) {
	projectAssignmentRepository := NewProjectAssignmentRepository(model.Db)
	err = projectAssignmentRepository.Save(p)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (p *ProjectAssignment) Update() (err error) {
	projectAssignmentRepository := NewProjectAssignmentRepository(model.Db)
	err = projectAssignmentRepository.Update(p)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (p *ProjectAssignment) Remove() (err error) {
	projectAssignmentRepository := NewProjectAssignmentRepository(model.Db)
	err = projectAssignmentRepository.Remove(p.ID)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

