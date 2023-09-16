package database

import (
    "log"
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
)

type Storage interface {
    CreateTask() error
    GetTasks() error
}

// some kind of configs
type Database struct {
    Connection *sql.DB
}

func NewSqliteStore() (*Database, error) {
    conn, err := sql.Open("sqlite3", "test.db")

    if err != nil {
        return nil, err
    }

    if err := conn.Ping(); err != nil {
        return nil, err
    }

    return &Database {
        Connection: conn,
    }, nil
}

func (db *Database) GetTasks() {
    sqlStmnt := `select * from todos;`

    execRes, err := db.Connection.Query(sqlStmnt)

    if err != nil {
        log.Printf("Error getting todos: %q: %s", err, sqlStmnt)
    }

    log.Println(execRes.Columns())

}

func (db *Database) CreateTask() {
    println("creating task!")

    	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	`
    _, err := db.Connection.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}



