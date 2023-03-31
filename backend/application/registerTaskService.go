package application

import (
	"log"

	"github.com/yukyoooo/go_next_ddd/domain/model/task"
	taskassignment "github.com/yukyoooo/go_next_ddd/domain/model/taskAssignment"
)

func RegisterTaskService(employeeId int, projectId int, milestoneId int, name string, detail string, status int, url string) error {
	newTask, err := task.NewTask(projectId, milestoneId, name, detail, status, url)
	if err != nil {
		return err
	}

	err = newTask.Create()
	if err != nil {
		return err
	}

	newTAssignment, err := taskassignment.NewTaskAssignment(newTask.ID, employeeId)
	if err != nil {
		return err
	}

	err = newTAssignment.Create()
	if err != nil {
		return err
	}

	log.Println("task is successfully created. task: ", newTask)
	return nil
}
