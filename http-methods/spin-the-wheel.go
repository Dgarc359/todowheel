package util

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	db "todowheel-backend/database"
	pb "todowheel-backend/proto"

	"google.golang.org/protobuf/proto"
)

func SpinTheWheel(w http.ResponseWriter, r *http.Request, conn *db.SqliteDatabase) {
	println("got get tasks request")
	p := &pb.PostGetTasks{}
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	proto.Unmarshal(body, p)

	res, _ := conn.GetTasks(p)

	randomTodoIndex := rand.Intn(len(res) - 0)
	randomTodo := res[randomTodoIndex]
	log.Printf("random item %v", randomTodo)

	w.Header().Set("Content-Type", "application/json")

	type SpinResponseStruct struct {
		Response
		Data db.UniqueTodo
	}

	json.NewEncoder(w).Encode(&SpinResponseStruct{Response{"success", 200}, randomTodo})
}
