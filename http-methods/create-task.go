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
	p := &pb.PostCreateTask{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("not able to read body")
	}
	r.Body.Close()

	proto.Unmarshal(body, p)

	fmt.Printf("Unmarshalled Body: %v", p)

	conn.CreateTask(p)

	json.NewEncoder(w).Encode(Response{"success", 200})
}
