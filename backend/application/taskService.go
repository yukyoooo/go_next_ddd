package application

import (
	"log"

	"github.com/yukyoooo/go_next_ddd/domain/model/task"
	taskassignment "github.com/yukyoooo/go_next_ddd/domain/model/taskAssignment"
)

type TaskApplicationService struct {
	taskRepository           task.TaskRepositorier
	taskAssignmentRepository taskassignment.TaskAssignmentRepositorier
}

func NewTaskApplicationService(taskRepository task.TaskRepositorier, taskAssignmentRepository taskassignment.TaskAssignmentRepositorier) *TaskApplicationService {
	return &TaskApplicationService{
		taskRepository:           taskRepository,
		taskAssignmentRepository: taskAssignmentRepository,
	}
}

func (tas *TaskApplicationService) Create(employeeId int, projectId int, milestoneId int, name string, detail string, status int, url string) error {
	newTask, err := task.NewTask(projectId, milestoneId, name, detail, status, url)
	if err != nil {
		return err
	}

	if err := tas.taskRepository.Save(newTask); err != nil {
		return err
	}

	taskId, err := tas.taskRepository.GetLastId()
	if err != nil {
		return err
	}

	newTAssignment, err := taskassignment.NewTaskAssignment(taskId, employeeId)
	if err != nil {
		return err
	}

	if err := tas.taskAssignmentRepository.Save(newTAssignment); err != nil {
		return err
	}

	log.Println("task is successfully created. task: ", newTask)
	return nil
}
