package projectassignment

type ProjectAssignmentRepository interface {
	Save(projectAssignment *ProjectAssignment) error
	Update(projectAssignment *ProjectAssignment) error
	Remove(id int) error
}
