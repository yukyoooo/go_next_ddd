package task

import (
	"log"

	"github.com/yukyoooo/go_next_ddd/domain/model"
)

type Task struct {
	ID          int
	projectID   int
	milestoneID int
	name        string
	detail      string
	status      int
	url         string
}

func (t *Task) GetID() int {
    return t.ID
}

func (t *Task) GetProjectID() int {
    return t.projectID
}

func (t *Task) GetMilestoneID() int {
    return t.milestoneID
}

func (t *Task) GetName() string {
    return t.name
}

func (t *Task) GetDetail() string {
    return t.detail
}

func (t *Task) GetStatus() int {
    return t.status
}

func (t *Task) Create() (err error){
    taskRepository := NewTaskRepository(model.Db)
	err = taskRepository.Save(t)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (t *Task) Update() (err error) {
	taskRepository := NewTaskRepository(model.Db)
	err = taskRepository.Update(t)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (t *Task) Remove() (err error) {
	taskRepository := NewTaskRepository(model.Db)
	err = taskRepository.Remove(t.ID)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func FindByID(id int) (*Task, error) {
	taskRepository := NewTaskRepository(model.Db)
	task, err := taskRepository.FindById(id)
	if err != nil {
		log.Fatal(err)
	}
	return task, nil
}

func FindByMilestoneId(milestoneID int) ([]*Task, error) {
	taskRepository := NewTaskRepository(model.Db)
	tasks, err := taskRepository.FindByMilestoneId(milestoneID)
	if err != nil {
		log.Fatal(err)
	}
	return tasks, nil
}

func FindByProjectId(projectID int) ([]*Task, error) {
	taskRepository := NewTaskRepository(model.Db)
	tasks, err := taskRepository.FindByProjectId(projectID)
	if err != nil {
		log.Fatal(err)
	}
	return tasks, nil
}