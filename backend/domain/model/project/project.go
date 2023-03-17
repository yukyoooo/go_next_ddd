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

func (p *Project) Update() (err error) {
	cmd := `update projects set name = ?, sort_id = ?, start_date = ?, end_date = ? where id = ?`
	_, err = model.Db.Exec(cmd, p.Name, p.SortID, p.StartDate, p.EndDate, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Project) Remove(id int) (err error) {
	cmd := `delete from projects where id = ?`
	_, err = model.Db.Exec(cmd, id)
	if err != nil {
		return err
	}
	return nil
}

func FindById(id int) (project *Project, err error) {
	cmd := `select id, name, sort_id, start_date, end_date from projects where id = ?`
	project = new(Project)
	err = model.Db.QueryRow(cmd, id).Scan(
		&project.ID,
		&project.Name,
		&project.SortID,
		&project.StartDate,
		&project.EndDate)
	if err != nil {
		return project, err
	}

	return project, nil
}