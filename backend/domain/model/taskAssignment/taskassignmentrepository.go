package taskassignment

import "database/sql"

type TaskAssignmentRepositorier interface {
	Save(taskAssignment *TaskAssignment) error
	Update(taskAssignment *TaskAssignment) error
	Remove(id int) error
}

type TaskAssignmentRepository struct {
	db *sql.DB
}

func NewTaskAssignmentRepository(db *sql.DB) (*TaskAssignmentRepository, error) {
	return &TaskAssignmentRepository{db: db}, nil
}

func (tar *TaskAssignmentRepository) Save(taskAssignment *TaskAssignment) error {
	tx, err := tar.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO task_assignments (
		task_id,
		employee_id) values (?, ?)`, taskAssignment.TaskID, taskAssignment.EmployeeID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (tar *TaskAssignmentRepository) Update(taskAssignment *TaskAssignment) error {
	_, err := tar.db.Exec(`UPDATE task_assignments SET task_id=?, employee_id=? WHERE id=?`, taskAssignment.TaskID, taskAssignment.EmployeeID, taskAssignment.ID)
	return err
}

func (tar *TaskAssignmentRepository) Remove(id int) error {
	_, err := tar.db.Exec("DELETE FROM task_assignments WHERE id=?", id)
	return err
}
