package projectassignment

import (
	"database/sql"
)

type ProjectAssignmentRepositorier interface {
	Save(projectAssignment *ProjectAssignment) error
	Update(projectAssignment *ProjectAssignment) error
	Remove(id int) error
}

type ProjectAssignmentRepository struct {
	db *sql.DB
}

func NewProjectAssignmentRepository(db *sql.DB) *ProjectAssignmentRepository {
	return &ProjectAssignmentRepository{db: db}
}

func (par *ProjectAssignmentRepository) Save(projectAssignment *ProjectAssignment) error {
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

func (par *ProjectAssignmentRepository) Update(projectAssignment *ProjectAssignment) error {
	_, err := par.db.Exec(`UPDATE project_assignments SET project_id=?, employee_id=? WHERE id=?`, projectAssignment.ProjectID, projectAssignment.EmployeeID, projectAssignment.ID)
	return err
}

func (par *ProjectAssignmentRepository) Remove(id int) error {
	_, err := par.db.Exec("DELETE FROM project_assignments WHERE id=?", id)
	return err
}