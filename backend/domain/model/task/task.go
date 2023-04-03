package task

type Task struct {
	ID          int
	ProjectID   int
	MilestoneID int
	Name        string
	Detail      string
	Status      int
	Url         string
}

func NewTask(projectID int, milestoneID int, name string, detail string, status int, url string) (*Task, error) {
	task := Task{
		ProjectID:   projectID,
		MilestoneID: milestoneID,
		Name:        name,
		Detail:      detail,
		Status:      status,
		Url:         url,
	}
	return &task, nil
}

func (t *Task) GetID() int {
	return t.ID
}

func (t *Task) GetProjectID() int {
	return t.ProjectID
}

func (t *Task) GetMilestoneID() int {
	return t.MilestoneID
}

func (t *Task) GetName() string {
	return t.Name
}

func (t *Task) GetDetail() string {
	return t.Detail
}

func (t *Task) GetStatus() int {
	return t.Status
}
