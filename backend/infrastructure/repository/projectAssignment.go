package repository

import (
	"database/sql"

	projectassignment "github.com/yukyoooo/go_next_ddd/domain/model/projectAssignment"
)

type ProjectAssignmentRepository struct {
	db *sql.DB
}

func NewProjectAssignmentRepository(db *sql.DB) (*ProjectAssignmentRepository, error) {
	return &ProjectAssignmentRepository{db: db}, nil
}

func (par *ProjectAssignmentRepository) Save(projectAssignment *projectassignment.ProjectAssignment) error {
	tx, err := par.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO project_assignments (
		project_id,
		employee_id) values (?, ?)`, projectAssignment.ProjectID, projectAssignment.EmployeeID)
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

func (par *ProjectAssignmentRepository) Update(projectAssignment *projectassignment.ProjectAssignment) error {
	_, err := par.db.Exec(`UPDATE project_assignments SET project_id=?, employee_id=? WHERE id=?`, projectAssignment.ProjectID, projectAssignment.EmployeeID, projectAssignment.ID)
	return err
}

func (par *ProjectAssignmentRepository) Remove(id int) error {
	_, err := par.db.Exec("DELETE FROM project_assignments WHERE id=?", id)
	return err
}
