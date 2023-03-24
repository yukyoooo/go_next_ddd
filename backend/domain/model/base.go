package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yukyoooo/go_next_ddd/config"
)

var Db *sql.DB

var err error

const (
	tableNameEmployee = "employees"
	tableNameProject = "projects"
	tableNameMilestone = "milestones"
	tableNameTask = "tasks"
	tableNameThread = "threads"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdE := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name STRING NOT NULL,
		last_name STRING NOT NULL,
		email STRING NOT NULL,
		password STRING NOT NULL,
		role INTEGER NOT NULL,
		created_at DATETIME TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
		updated_at DATETIME TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
		)`, tableNameEmployee)
	Db.Exec(cmdE)

	cmdP := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        name STRING NOT NULL,
        sort_id INTEGER NOT NULL,
        start_date DATETIME NOT NULL,
        end_date DATETIME NOT NULL,
		created_at DATETIME TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
		updated_at DATETIME TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
		)`, tableNameProject)
	Db.Exec(cmdP)

	cmdM := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        name STRING NOT NULL,
        start_date DATETIME	NOT NULL,
        end_date DATETIME NOT NULL,
		created_at DATETIME TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
		updated_at DATETIME TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
        )`, tableNameMilestone)
    Db.Exec(cmdM)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		project_id INTEGER NOT NULL,
		milestone_id INTEGER NOT NULL,
        name STRING NOT NULL,
		detail STRING NOT NULL,
		status INTEGER NOT NULL,
		url STRING NOT NULL,
		created_at DATETIME TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
		updated_at DATETIME TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
        )`, tableNameTask)
	Db.Exec(cmdT)

	cmdTh := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_id INTEGER NOT NULL,
		from_employee_id INTEGER NOT NULL,
		to_employee_id INTEGER NOT NULL,
		title STRING NOT NULL,
		body STRING, 
		resolution_flag BOOLEAN NOT NULL,
		created_at DATETIME TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
		updated_at DATETIME TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
		)`, tableNameThread)
	Db.Exec(cmdTh)
}
