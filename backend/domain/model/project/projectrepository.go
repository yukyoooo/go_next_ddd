package project

type ProjectRepository interface {
	Save(project *Project) error
	FindById(id int) (*Project, error)
	Update(project *Project) error
	Remove(id int) error
	GetLastSortId() (int, error)
	GetLastId() (int, error)
}
