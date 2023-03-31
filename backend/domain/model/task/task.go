package task

type Task struct {
	ID          int
	projectID   int
	milestoneID int
	name        string
	detail      string
	status      int
	url         string
}

func NewTask(projectID int, milestoneID int, name string, detail string, status int, url string) (*Task, error) {
	task := Task{
		projectID:   projectID,
		milestoneID: milestoneID,
		name:        name,
		detail:      detail,
		status:      status,
		url:         url,
	}
	return &task, nil
}

func (t *Task) GetID() int {
	return t.ID
}

func (t *Task) GetProjectID() int {
	return t.projectID
}

func (t *Task) GetMilestoneID() int {
	return t.milestoneID
}

func (t *Task) GetName() string {
	return t.name
}

func (t *Task) GetDetail() string {
	return t.detail
}

func (t *Task) GetStatus() int {
	return t.status
}
