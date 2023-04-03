package application

import (
	"log"
	"time"

	"github.com/yukyoooo/go_next_ddd/domain/model/project"
	projectassignment "github.com/yukyoooo/go_next_ddd/domain/model/projectAssignment"
)

type ProjectApplicationService struct {
	projectRepository           project.ProjectRepository
	projectAssignmentRepository projectassignment.ProjectAssignmentRepositorier
}

func NewProjectApplicationService(projectRepository project.ProjectRepository, projectAssignmentRepository projectassignment.ProjectAssignmentRepositorier) *ProjectApplicationService {
	return &ProjectApplicationService{
		projectRepository:           projectRepository,
		projectAssignmentRepository: projectAssignmentRepository,
	}
}

func (pas *ProjectApplicationService) Create(employeeId int, name string, startDate time.Time, EndDate time.Time) error {
	sortId, err := pas.projectRepository.GetLastSortId()
	if err != nil {
		return err
	}

	newProject, err := project.NewProject(sortId, name, startDate, EndDate)
	if err != nil {
		return err
	}

	if err := pas.projectRepository.Save(newProject); err != nil {
		return err
	}

	projectId, err := pas.projectRepository.GetLastId()
	if err != nil {
		return err
	}

	newPAssignment, err := projectassignment.NewProjectAssignment(projectId, employeeId)
	if err != nil {
		return err
	}

	if err := pas.projectAssignmentRepository.Save(newPAssignment); err != nil {
		return err
	}

	log.Println("project is successfully created. project: ", newProject)
	return nil
}
