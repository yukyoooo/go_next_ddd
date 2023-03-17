package project

import (
	"log"
	"time"

	model "github.com/yukyoooo/go_next_ddd/domain/model"
)

type Project struct {
	ID int
	Name string
	SortID int
	StartDate time.Time
	EndDate time.Time
}

func NewProject(name string, sortId int, startDate time.Time, EndDate time.Time) (*Project, error) {
	return &Project{Name: name, SortID: sortId, StartDate: startDate, EndDate: EndDate}, nil
}

func (p *Project) Create() (err error) {
	projectRepository := NewProjectRepository(model.Db)
	err = projectRepository.Save(p)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (p *Project) Update() (err error) {
	projectRepository := NewProjectRepository(model.Db)
    err = projectRepository.Update(p)
    if err!= nil {
        log.Fatalln(err)
    }
    return err
}

func (p *Project) Remove(id int) (err error) {
	projectRepository := NewProjectRepository(model.Db)
	err = projectRepository.Remove(id)
	if err!= nil {
        log.Fatalln(err)
    }
	return err
}

func FindById(id int) (project *Project, err error) {
	projectRepository := NewProjectRepository(model.Db)
	project, err = projectRepository.FindById(id)
	if err!= nil {
        log.Fatalln(err)
    }
	return project, err
}