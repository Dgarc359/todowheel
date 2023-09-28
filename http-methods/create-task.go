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

func CreateTask(w http.ResponseWriter, r *http.Request, conn *db.SqliteDatabase) {
	println("Got create task request")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("not able to read body")
	}
	r.Body.Close()

	applicationType := r.Header.Get("Content-Type")

	p := &pb.PostCreateTask{}
	if applicationType == "application/proto" {
		proto.Unmarshal(body, p)
	} else if applicationType == "application/json" {
		if err := json.Unmarshal(body, &p); err != nil {
			json.NewEncoder(w).Encode(Response{"bad request", 400})
			return
		}

	}

	fmt.Printf("Unmarshalled Body: %v", p)

	conn.CreateTask(p)

	json.NewEncoder(w).Encode(Response{"success", 200})
}
