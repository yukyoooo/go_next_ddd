package application

import (
	"log"
	"time"

	"github.com/yukyoooo/go_next_ddd/domain/model/milestone"
)

func RegisterMilestoneService(projectId int, name string, startDate time.Time, EndDate time.Time) (error){
	newMilestone, err := milestone.NewMilestone(projectId, name, startDate, EndDate)
	if err != nil {
		return err
	}

	err = newMilestone.Create()
	if err != nil {
		return err 
	}

	log.Println("milestone is successfully created. milestone: ", newMilestone)
	return nil
}