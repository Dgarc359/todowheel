package util

import (
	"fmt"
	"io"
	"net/http"
	pb "todowheel-backend/proto"

	"google.golang.org/protobuf/proto"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	println("Got create task request")
	p := &pb.PostGetTask{}
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	proto.Unmarshal(body, p)

	fmt.Printf("Unmarshalled Body: %v", p)
}
