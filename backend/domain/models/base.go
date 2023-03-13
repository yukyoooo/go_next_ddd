package models

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
}
