package util

import (
	"io"
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

	conn.GetTasks(p)
}
