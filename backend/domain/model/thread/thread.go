package thread

import (
	"time"
)

type Thread struct {
	ID             int
	TaskID         int
	FromEmployeeID int
	ToEmployeeID   int
	Title          string
	Body           string
	ResolutionFlag int
	CreatedAt      time.Time
}

func NewThread(taskID int, fromEmployeeID int, toEmployeeID int, title string, body string, resolutionFlag int) *Thread {
	return &Thread{
		TaskID:         taskID,
		FromEmployeeID: fromEmployeeID,
		ToEmployeeID:   toEmployeeID,
		Title:          title,
		Body:           body,
		ResolutionFlag: resolutionFlag,
		CreatedAt:      time.Now(),
	}
}
