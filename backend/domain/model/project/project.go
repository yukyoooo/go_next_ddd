package project

import (
	"time"
)

type Project struct {
	ID        int
	Name      string
	SortID    int
	StartDate time.Time
	EndDate   time.Time
}

func NewProject(sortId int, name string, startDate time.Time, EndDate time.Time) (*Project, error) {
	return &Project{Name: name, SortID: sortId + 1, StartDate: startDate, EndDate: EndDate}, nil
}
