package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	db "todowheel-backend/database"
	pb "todowheel-backend/proto"

	"google.golang.org/protobuf/proto"
)

func GetTasks(w http.ResponseWriter, r *http.Request, conn *db.SqliteDatabase) {
	println("got get tasks request")
	p := &pb.PostGetTasks{}
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	proto.Unmarshal(body, p)

	res, err := conn.GetTasks(p)

	if err != nil {
		log.Printf("Failed to get tasks")
		return
	}
	w.Header().Set("Content-Type", "application/json")

	type GetTasksResponse struct {
		Response
		Data []db.UniqueTodo
	}

	json.NewEncoder(w).Encode(&GetTasksResponse{Response{"success", 200}, res})

}
