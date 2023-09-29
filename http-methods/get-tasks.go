package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	db "todowheel-backend/database"
	pb "todowheel-backend/proto"

	"google.golang.org/protobuf/proto"
)

func GetTasks(w http.ResponseWriter, r *http.Request, conn *db.SqliteDatabase) {
	println("got get tasks request")
	p := &pb.PostGetTasks{}
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	fmt.Printf("Body: %+v", string(body))
	applicationType := r.Header.Get("Content-Type")

	if applicationType == "application/proto" {
		proto.Unmarshal(body, p)
	} else if applicationType == "application/json" {
		if err := json.Unmarshal(body, &p); err != nil {
			json.NewEncoder(w).Encode(Response{"bad request", 400})
			return
		}
	}

	fmt.Printf("Calling get tasks: %+v", p)
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
