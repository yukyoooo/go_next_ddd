package project

import "database/sql"

type ProjectRepositorier interface {
	Save(project *Project) (error)
	FindById(id int) (*Project, error)
	Update(project *Project) (error)
	Remove(id int) (error)
}

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
    return &ProjectRepository{db: db}
}

func (pr *ProjectRepository) Save(project *Project) (*Project, error) {
    tx, err := pr.db.Begin()
    if err != nil {
        return nil, err
    }

    _, err = tx.Exec(`INSERT INTO projects (
        name,
        sort_id,
        start_date,
        end_date) values (?, ?, ?, ?)`, project.Name, project.SortID, project.StartDate, project.EndDate)
    if err != nil {
        tx.Rollback()
        return nil, err
    }

    err = tx.Commit()
    if err != nil {
        return nil, err
    }

    err = pr.db.QueryRow("SELECT id FROM projects order by id desc limit 1").Scan(
        &project.ID)
    if err != nil {
        return nil, err
    }
    
    return project, nil
}

func (pr *ProjectRepository) FindById(id int) (*Project, error) {
	var project Project
    err := pr.db.QueryRow("SELECT * FROM projects WHERE id =?", id).Scan(
		&project.ID,
		&project.Name,
        &project.SortID,
        &project.StartDate,
        &project.EndDate)
	if err!= nil {
        return nil, err
    }
	return &project, nil
}

func (pr *ProjectRepository) Update(project *Project) (error) {
	_, err := pr.db.Exec("UPDATE projects SET name =?, sort_id =?, start_date =?, end_date =? WHERE id =?",
        project.Name,
        project.SortID,
        project.StartDate,
        project.EndDate,
        project.ID)
    return err
}

func (pr *ProjectRepository) Remove(id int) (error) {
	_, err := pr.db.Exec("DELETE FROM projects WHERE id =?", id)
    return err
}