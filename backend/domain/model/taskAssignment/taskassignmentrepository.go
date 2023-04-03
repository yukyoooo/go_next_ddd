package taskassignment

type TaskAssignmentRepository interface {
	Save(taskAssignment *TaskAssignment) error
	Update(taskAssignment *TaskAssignment) error
	Remove(id int) error
}
