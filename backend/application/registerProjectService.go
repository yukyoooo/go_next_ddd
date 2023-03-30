package application

import (
	"log"
	"time"

	"github.com/yukyoooo/go_next_ddd/domain/model/project"
	projectassignment "github.com/yukyoooo/go_next_ddd/domain/model/projectAssignment"
)

func RegisterProjectService(employeeId int, name string, startDate time.Time, EndDate time.Time) (error){
	newProject, err := project.NewProject(name, startDate, EndDate)
	if err != nil {
		return err
	}

	newProject, err = newProject.Create()
	if err != nil {
		return err 
	}

	newPAssignment, err := projectassignment.NewProjectAssignment(newProject.ID, employeeId)
	if err != nil {
		return err
	}


	err = newPAssignment.Create()
	if err != nil {
		return err
	}

	log.Println("project is successfully created. project: ", newProject)
	return nil
}