package thread

type ThreadRepository interface {
	Save(thread *Thread) error
	FindById(id int) (*Thread, error)
	FindByTaskId(taskId int) ([]*Thread, error)
	Update(thread *Thread) error
	Remove(id int) error
}
