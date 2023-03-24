package thread

import (
	"log"
	"time"

	"github.com/yukyoooo/go_next_ddd/domain/model"
)

type Thread struct {
	ID             int
	taskID         int
	fromEmployeeID int
	toEmployeeID   int
	title          string
	body           string
	resolutionFlag int
	createdAt 	   time.Time
}

func NewThread(taskID int, fromEmployeeID int, toEmployeeID int, title string, body string, resolutionFlag int) *Thread {
	return &Thread{
		taskID:         taskID,
		fromEmployeeID: fromEmployeeID,
		toEmployeeID:   toEmployeeID,
		title:          title,
		body:           body,
		resolutionFlag: resolutionFlag,
		createdAt:      time.Now(),
	}
}

func (t *Thread) Create() (err error) {
	threadRepository := NewThreadRepository(model.Db)
	err = threadRepository.Save(t)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (t *Thread) Update() (err error) {
	threadRepository := NewThreadRepository(model.Db)
	err = threadRepository.Update(t)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (t *Thread) Remove() (err error) {
	threadRepository := NewThreadRepository(model.Db)
	err = threadRepository.Remove(t.ID)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func FindByID(id int) (thread *Thread, err error) {
	threadRepository := NewThreadRepository(model.Db)
	thread, err = threadRepository.FindById(id)
	if err != nil {
		log.Fatal(err)
	}
	return thread, nil
}

func FindByTaskID(taskID int) (threads []*Thread, err error) {
	threadRepository := NewThreadRepository(model.Db)
	threads, err = threadRepository.FindByTaskId(taskID)
	if err != nil {
		log.Fatal(err)
	}
	return threads, nil
}