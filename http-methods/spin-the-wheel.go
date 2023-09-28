package util

import (
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
	log.Printf("random item %v", res[randomTodoIndex])

}
