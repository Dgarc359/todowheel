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
	fmt.Println("got get task request")
	p := &pb.PostGetTask{}
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	applicationType := r.Header.Get("Content-Type")

	var data map[string]interface{}

	if applicationType == "application/proto" {
		proto.Unmarshal(body, p)
	} else if applicationType == "application/json" {
		if err := json.Unmarshal(body, &data); err != nil {
			json.NewEncoder(w).Encode(Response{"bad request", 400})
			return
		}
		p.TaskName = data["taskName"].(string)
	}

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
