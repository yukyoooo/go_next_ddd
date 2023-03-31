package milestone

import (
	"database/sql"

	"github.com/yukyoooo/go_next_ddd/domain/model"
)

type MilestoneRepositorier interface {
	Save(milestone *Milestone) error
	FindById(id int) (*Milestone, error)
	Update(milestone *Milestone) error
	Remove(id int) error
}

type MilestoneRepository struct {
	db *sql.DB
}

func NewMilestoneRepository(db *sql.DB) (*MilestoneRepository, error) {
	return &MilestoneRepository{db}, nil
}

func (mr *MilestoneRepository) Save(milestone *Milestone) error {
	tx, err := mr.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO milestones (
		project_id,
        name,
		start_date,
		end_date) values (?, ?, ?, ?)`, milestone.ProjectID, milestone.Name, milestone.StartDate.Format("2006-01-02 15:04:05"), milestone.EndDate.Format("2006-01-02 15:04:05"))
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

func (mr *MilestoneRepository) FindById(id int) (*Milestone, error) {
	var milestone Milestone
	err := model.Db.QueryRow("SELECT id, project_id, name, start_date, end_date FROM milestones WHERE id=?", id).Scan(
		&milestone.ID,
		&milestone.ProjectID,
		&milestone.Name,
		&milestone.StartDate,
		&milestone.EndDate)
	if err != nil {
		return nil, err
	}
	return &milestone, nil
}

func (mr *MilestoneRepository) Update(milestone *Milestone) error {
	_, err := mr.db.Exec("UPDATE milestones SET project_id=?, name=?, start_date=?, end_date=? WHERE id=?", milestone.ProjectID, milestone.Name, milestone.StartDate.Format("2006-01-02 15:04:05"), milestone.EndDate.Format("2006-01-02 15:04:05"), milestone.ID)
	return err
}

func (mr *MilestoneRepository) Remove(id int) error {
	_, err := mr.db.Exec("DELETE FROM milestones WHERE id=?", id)
	return err
}
