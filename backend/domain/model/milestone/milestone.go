package milestone

import (
	"log"
	"time"

	"github.com/yukyoooo/go_next_ddd/domain/model"
	ierrors "github.com/yukyoooo/go_next_ddd/ierrors"
)

type Milestone struct {
	ID          int
	Name        string
	StartDate   time.Time
	EndDate     time.Time
}

func NewMilestone(name string, startDate time.Time, endDate time.Time) (*Milestone, error) {
    if name == "" {
        return nil, ierrors.ErrInvalidName
    }
	
    if startDate.IsZero() {
        return nil, ierrors.ErrInvalidDate
    }

    if endDate.IsZero() {
        return nil, ierrors.ErrInvalidDate
    }

    return &Milestone{
        Name:        name,
        StartDate:   startDate,
        EndDate:     endDate,
    }, nil
}

func (m *Milestone) GetID() int {
    return m.ID
}

func (m *Milestone) GetName() string {
    return m.Name
}

func (m *Milestone) GetStartDate() time.Time {
    return m.StartDate
}

func (m *Milestone) GetEndDate() time.Time {
    return m.EndDate
}

func (m *Milestone) IsStartBeforeEnd() bool {
    return m.StartDate.Before(m.EndDate)
}

func (m *Milestone) IsStartAfterEnd() bool {
    return m.StartDate.After(m.EndDate)
}

func (m *Milestone) Create() (err error) {
	milestoneRepository := NewMilestoneRepository(model.Db)
	err = milestoneRepository.Save(m)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (m *Milestone) Update() (err error) {
	milestoneRepository := NewMilestoneRepository(model.Db)
	err = milestoneRepository.Update(m)
	if err!= nil {
        log.Fatal(err)
    }
	return nil
}

func (m *Milestone) Remove() (err error) {
	milestoneRepository := NewMilestoneRepository(model.Db)
    err = milestoneRepository.Remove(m.ID)
    if err!= nil {
        log.Fatal(err)
    }
    return nil
}

func FindByID(id int) (milestone *Milestone, err error) {
	milestoneRepository := NewMilestoneRepository(model.Db)
    milestone, err = milestoneRepository.FindById(id)
    if err!= nil {
        log.Fatal(err)
    }
    return milestone, nil
}