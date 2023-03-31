package milestone

import (
	"time"

	ierrors "github.com/yukyoooo/go_next_ddd/ierrors"
)

type Milestone struct {
	ID        int
	ProjectID int
	Name      string
	StartDate time.Time
	EndDate   time.Time
}

func NewMilestone(projectId int, name string, startDate time.Time, endDate time.Time) (*Milestone, error) {
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
		ProjectID: projectId,
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
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
