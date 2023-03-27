package taskassignment

import (
	"log"

	"github.com/yukyoooo/go_next_ddd/domain/model"
)

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

func (t *TaskAssignment) Create() (err error) {
	taskAssignmentRepository := NewTaskAssignmentRepository(model.Db)
	err = taskAssignmentRepository.Save(t)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (t *TaskAssignment) Update() (err error) {
	taskAssignmentRepository := NewTaskAssignmentRepository(model.Db)
	err = taskAssignmentRepository.Update(t)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (t *TaskAssignment) Remove() (err error) {
	taskAssignmentRepository := NewTaskAssignmentRepository(model.Db)
	err = taskAssignmentRepository.Remove(t.ID)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}