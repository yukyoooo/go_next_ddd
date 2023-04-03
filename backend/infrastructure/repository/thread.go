package repository

import (
	"database/sql"

	"github.com/yukyoooo/go_next_ddd/domain/model"
	"github.com/yukyoooo/go_next_ddd/domain/model/thread"
)

type ThreadRepository struct {
	db *sql.DB
}

func NewThreadRepository(db *sql.DB) *ThreadRepository {
	return &ThreadRepository{db}
}

func (tr *ThreadRepository) Save(thread *thread.Thread) error {
	tx, err := tr.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO threads (
		task_id,
		from_employee_id,
		to_employee_id,
		title,
		body,
		resolution_flag) values (?, ?, ?, ?, ?, ?)`, thread.TaskID, thread.FromEmployeeID, thread.ToEmployeeID, thread.Title, thread.Body, thread.ResolutionFlag)
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

func (tr *ThreadRepository) FindById(id int) (*thread.Thread, error) {
	var thread thread.Thread
	err := model.Db.QueryRow("SELECT id, task_id, from_employee_id, to_employee_id, title, body, resolution_flag, created_at  FROM threads WHERE id=?", id).Scan(
		&thread.ID,
		&thread.TaskID,
		&thread.FromEmployeeID,
		&thread.ToEmployeeID,
		&thread.Title,
		&thread.Body,
		&thread.ResolutionFlag,
		&thread.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &thread, nil
}

func (tr *ThreadRepository) FindByTaskId(taskId int) ([]*thread.Thread, error) {
	rows, err := model.Db.Query("SELECT id, task_id, from_employee_id, to_employee_id, title, body, resolution_flag, created_at  FROM threads WHERE task_id=?", taskId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []*thread.Thread
	for rows.Next() {
		var thread thread.Thread
		err := rows.Scan(
			&thread.ID,
			&thread.TaskID,
			&thread.FromEmployeeID,
			&thread.ToEmployeeID,
			&thread.Title,
			&thread.Body,
			&thread.ResolutionFlag,
			&thread.CreatedAt)
		if err != nil {
			return nil, err
		}
		threads = append(threads, &thread)
	}
	return threads, nil
}

func (tr *ThreadRepository) Update(thread *thread.Thread) error {
	tx, err := tr.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE threads SET
		task_id = ?,
		from_employee_id = ?,
		to_employee_id = ?,
		title = ?,
		body = ?,
		resolution_flag = ? WHERE id = ?`, thread.TaskID, thread.FromEmployeeID, thread.ToEmployeeID, thread.Title, thread.Body, thread.ResolutionFlag, thread.ID)
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

func (tr *ThreadRepository) Remove(id int) error {
	_, err := model.Db.Exec("DELETE FROM threads WHERE id=?", id)
	return err
}
