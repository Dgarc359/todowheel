package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	db "todowheel-backend/database"
	pb "todowheel-backend/proto"

	"google.golang.org/protobuf/proto"
)

func SpinTheWheel(w http.ResponseWriter, r *http.Request, conn *db.SqliteDatabase) {
	fmt.Println("got spin the wheel request")
	p := &pb.PostGetTasks{}
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	proto.Unmarshal(body, p)

	res, _ := conn.GetTasks(p)

	randomTodoIndex := rand.Intn(len(res))
	randomTodo := res[randomTodoIndex]
	log.Printf("random item %v", randomTodo)

	w.Header().Set("Content-Type", "application/json")

	type SpinResponseStruct struct {
		Response
		Data db.UniqueTodo
	}

	json.NewEncoder(w).Encode(&SpinResponseStruct{Response{"success", 200}, randomTodo})
}
