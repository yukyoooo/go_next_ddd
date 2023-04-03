package milestone

type MilestoneRepository interface {
	Save(milestone *Milestone) error
	FindById(id int) (*Milestone, error)
	Update(milestone *Milestone) error
	Remove(id int) error
}
