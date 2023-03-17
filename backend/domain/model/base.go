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
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdE := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name STRING,
		last_name STRING,
		email STRING,
		password STRING,
		role INTEGER
		)`, tableNameEmployee)
	Db.Exec(cmdE)

	cmdP := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        name STRING,
        sort_id INTEGER,
        start_date DATETIME,
        end_date DATETIME
		)`, tableNameProject)
	Db.Exec(cmdP)

	cmdM := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        name STRING,
        start_date DATETIME,
        end_date DATETIME
        )`, tableNameMilestone)
    Db.Exec(cmdM)
}
