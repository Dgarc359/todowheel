package database

import (
    "log"
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
)

// some kind of configs
type Database struct {
    Connection *sql.DB
}

func (db *Database) New() (*Database, error) {
    conn, err := sql.Open("sqlite3", "./test.db")

    if err != nil {
        return &Database{}, err
    }

    return &Database {
        Connection: conn,
    }, nil
}

func (db *Database) CreateTask() {
    println("creating task!")

    	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
    _, err := db.Connection.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}



