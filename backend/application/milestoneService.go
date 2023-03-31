package application

import (
	"log"
	"time"

	"github.com/yukyoooo/go_next_ddd/domain/model/milestone"
)

type MilestoneApplicationService struct {
	milestoneRepository milestone.MilestoneRepositorier
}

func NewMilestoneApplicationService(milestoneRepository milestone.MilestoneRepositorier) *MilestoneApplicationService {
	return &MilestoneApplicationService{
		milestoneRepository: milestoneRepository,
	}
}

func (mas *MilestoneApplicationService) Create(projectId int, name string, startDate time.Time, EndDate time.Time) error {
	newMilestone, err := milestone.NewMilestone(projectId, name, startDate, EndDate)
	if err != nil {
		return err
	}

	if err := mas.milestoneRepository.Save(newMilestone); err != nil {
		return err
	}

	log.Println("milestone is successfully created. milestone: ", newMilestone)
	return nil
}
