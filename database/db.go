package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
    pb "todowheel-backend/proto"
)

type Todo struct {
    ID string `gorm:"column:id" json:"id"`
    TaskName string `gorm:"column:todo_name" json:"taskName"`
    TaskLength int64 `gorm:"column:todo_length" json:"taskLength"`
}


// some kind of configs
type SqliteDatabase struct {
    Connection *sql.DB
}

func NewSqliteStore() (*SqliteDatabase, error) {
    conn, err := sql.Open("sqlite3", "test.db")

    if err != nil {
        return nil, err
    }

    if err := conn.Ping(); err != nil {
        return nil, err
    }

    return &SqliteDatabase {
        Connection: conn,
    }, nil
}

func (db *SqliteDatabase) GetTasks() {
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

func (db *SqliteDatabase) CreateTask(t *pb.PostCreateTask) {
    println("creating task!")

    queryString := `insert into todos
    (todo_name, todo_length)
    values
    (?, ?)`

    preparedQuery, err := db.Connection.Prepare(queryString)

    if err != nil {
        log.Fatal("failed to prepare query")
    }

    _,err = preparedQuery.Exec(t.TaskName, t.TaskLength)

    if err != nil {
        log.Fatal(err)
    }
}



