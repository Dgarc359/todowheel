package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	db "todowheel-backend/database"
	pb "todowheel-backend/proto"

	"google.golang.org/protobuf/proto"
)

func GetTask(w http.ResponseWriter, r *http.Request, conn *db.SqliteDatabase) {
	p := &pb.PostGetTask{}
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	proto.Unmarshal(body, p)

	fmt.Printf("Unmarshalled Body: %v", p)

	res, err := conn.GetTask(p)
	if err != nil {
		log.Printf("Failed to get tasks")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	type GetTaskResponse struct {
		Response
		Data db.UniqueTodo
	}

	json.NewEncoder(w).Encode(&GetTaskResponse{Response{"success", 200}, res})

}
