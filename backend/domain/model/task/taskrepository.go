package task

import (
	"database/sql"
)

type TaskRepositorier interface {
	Save(task *Task) error
	FindByMilestoneId(milestoneId int) ([]*Task, error)
	FindByProjectId(projectId int) ([]*Task, error)
	FindById(id int) (*Task, error)
	Update(task *Task) error
	Remove(id int) error
}

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (tr *TaskRepository) Save(task *Task) error {
	tx, err := tr.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO tasks (
		project_id,
		milestone_id,
		name,
		detail,
		status,
		url) values (?, ?, ?, ?, ?, ?)`, task.projectID, task.milestoneID, task.name, task.detail, task.status, task.url)
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

func (tr *TaskRepository) FindByMilestoneId(milestoneId int) ([]*Task, error) {
	rows, err := tr.db.Query("SELECT id, project_id, milestone_id, name, detail, status, url FROM tasks WHERE milestone_id=?", milestoneId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.projectID, &task.milestoneID, &task.name, &task.detail, &task.status, &task.url); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (tr *TaskRepository) FindByProjectId(projectId int) ([]*Task, error) {
	rows, err := tr.db.Query("SELECT id, project_id, milestone_id, name, detail, status FROM tasks WHERE project_id=?", projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.projectID, &task.milestoneID, &task.name, &task.detail, &task.status); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (tr *TaskRepository) FindById(id int) (*Task, error) {
	var task Task
	err := tr.db.QueryRow("SELECT id, project_id, milestone_id, name, detail, status, url FROM tasks WHERE id=?", id).Scan(
		&task.ID,
		&task.projectID,
		&task.milestoneID,
		&task.name,
		&task.detail,
		&task.status)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (tr *TaskRepository) Update(task *Task) error {
	_, err := tr.db.Exec(`UPDATE tasks SET project_id=?, milestone_id=?, name=?, detail=?, status=?, url=?, WHERE id=?`, task.projectID, task.milestoneID, task.name, task.detail, task.status, task.url, task.ID)
	return err
}

func (tr *TaskRepository) Remove(id int) error {
	_, err := tr.db.Exec("DELETE FROM tasks WHERE id=?", id)
	return err
}