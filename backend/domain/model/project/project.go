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

func (p *Project) CreateProject() (err error) {
	cmd := `insert into projects (
		name,
		sort_id,
        start_date,
        end_date
	) values (?, ?, ?, ?)`

	_, err = model.Db.Exec(cmd, p.Name, p.SortID, p.StartDate, p.EndDate)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}