package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	pb "todowheel-backend/proto"
)

type UniqueTodo struct {
	ID         string `gorm:"column:id" json:"id"`
	TaskName   string `gorm:"column:todo_name" json:"taskName"`
	TaskLength int64  `gorm:"column:todo_length" json:"taskLength"`
}

type GetTasks struct {
	MinTaskLength int32 `gorm:"column:todo_length" json:"taskLength"`
	MaxTaskLength int32 `gorm:"column:todo_length" json:"taskLength"`
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

	return &SqliteDatabase{
		Connection: conn,
	}, nil
}

func (db *SqliteDatabase) GetTasks(t *pb.PostGetTasks) ([]UniqueTodo, error) {
	query := ""

	if t.MaxTaskLength == nil && t.MinTaskLength == nil {
		query = `select * from todos;`
		execRes, err := db.Connection.Query(query)
		if err != nil {
			log.Printf("Error getting todos: %q: %s", err, query)
		}

		for cont := true; cont; cont = execRes.NextResultSet() {
			for execRes.Next() {
				var t UniqueTodo
				err := execRes.Scan(&t.ID, &t.TaskName, &t.TaskLength)
				if err != nil {
					log.Fatal("error getting row")
				}
				fmt.Printf("%+v\n", t)
				log.Println(t)
			}
		}
	}
	query = `
        select * from todos
        where todo_length > ?
        and todo_length < ?
    `

	preparedQuery, err := db.Connection.Prepare(query)
	if err != nil {
		log.Printf("error preparing query: %q", err)
	}

	getTasks := &GetTasks{
		MinTaskLength: *t.MinTaskLength,
		MaxTaskLength: *t.MaxTaskLength,
	}

	rows, err := preparedQuery.Query(getTasks.MinTaskLength, getTasks.MaxTaskLength)
	defer rows.Close()

	if err != nil {
		log.Printf("Error getting todos: %q: %s", err, query)
	}

	var todos []UniqueTodo
	for rows.Next() {
		var uTodo UniqueTodo
		rows.Scan(&uTodo.ID, &uTodo.TaskName, &uTodo.TaskLength)
		log.Printf("Row: %v", uTodo)
		todos = append(todos, uTodo)
	}
	return todos, nil
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

	_, err = preparedQuery.Exec(t.TaskName, t.TaskLength)

	if err != nil {
		log.Fatal(err)
	}

	println("made entry in db")
}
