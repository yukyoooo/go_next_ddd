package repository

import (
	"database/sql"

	"github.com/yukyoooo/go_next_ddd/domain/model/task"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) (*TaskRepository, error) {
	return &TaskRepository{db}, nil
}

func (tr *TaskRepository) Save(task *task.Task) error {
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
		url) values (?, ?, ?, ?, ?, ?)`, task.ProjectID, task.MilestoneID, task.Name, task.Detail, task.Status, task.Url)
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

func (tr *TaskRepository) FindByMilestoneId(milestoneId int) ([]*task.Task, error) {
	rows, err := tr.db.Query("SELECT id, project_id, milestone_id, name, detail, status, url FROM tasks WHERE milestone_id=?", milestoneId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*task.Task
	for rows.Next() {
		var task task.Task
		if err := rows.Scan(&task.ID, &task.ProjectID, &task.MilestoneID, &task.Name, &task.Detail, &task.Status, &task.Url); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (tr *TaskRepository) FindByProjectId(projectId int) ([]*task.Task, error) {
	rows, err := tr.db.Query("SELECT id, project_id, milestone_id, name, detail, status FROM tasks WHERE project_id=?", projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*task.Task
	for rows.Next() {
		var task task.Task
		if err := rows.Scan(&task.ID, &task.ProjectID, &task.MilestoneID, &task.Name, &task.Detail, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (tr *TaskRepository) FindById(id int) (*task.Task, error) {
	var task task.Task
	err := tr.db.QueryRow("SELECT id, project_id, milestone_id, name, detail, status, url FROM tasks WHERE id=?", id).Scan(
		&task.ID,
		&task.ProjectID,
		&task.MilestoneID,
		&task.Name,
		&task.Detail,
		&task.Status)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (tr *TaskRepository) Update(task *task.Task) error {
	_, err := tr.db.Exec(`UPDATE tasks SET project_id=?, milestone_id=?, name=?, detail=?, status=?, url=?, WHERE id=?`, task.ProjectID, task.MilestoneID, task.Name, task.Detail, task.Status, task.Url, task.ID)
	return err
}

func (tr *TaskRepository) Remove(id int) error {
	_, err := tr.db.Exec("DELETE FROM tasks WHERE id=?", id)
	return err
}

func (tr *TaskRepository) GetLastId() (int, error) {
	var id int
	err := tr.db.QueryRow("SELECT id FROM tasks ORDER BY id DESC LIMIT 1").Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 1, nil
		} else {
			return 1, err
		}
	}
	return id, nil
}
