package task

type TaskRepository interface {
	Save(task *Task) error
	FindByMilestoneId(milestoneId int) ([]*Task, error)
	FindByProjectId(projectId int) ([]*Task, error)
	FindById(id int) (*Task, error)
	Update(task *Task) error
	Remove(id int) error
	GetLastId() (int, error)
}
