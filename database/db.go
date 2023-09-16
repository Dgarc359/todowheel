package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	_ "gorm.io/gorm"
)

type Todo struct {
    ID string `gorm:"column:id" json:"id"`
    TaskName string `gorm:"column:todo_name" json:"taskName"`
    TaskLength string `gorm:"column:todo_length" json:"taskLength"`
}

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


    log.Println("getting rows")
    for cont := true; cont; cont = execRes.NextResultSet() {
        for execRes.Next() {
            var t Todo
            err := execRes.Scan(&t.ID, &t.TaskName, &t.TaskLength)
            if err != nil {
                log.Fatal("error getting row")
            }
            fmt.Printf("%+v\n", t)
            log.Println(t)
        }
    }
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



