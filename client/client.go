package client

import (
	"bytes"
	"log"
	"net/http"
	pb "todowheel-backend/proto"

	"google.golang.org/protobuf/proto"
)

func CreateClient() {
    var temp int32 = 20
    createTask := new(pb.PostCreateTask)
    createTask.TaskName = "test-task"
    createTask.TaskId = 3
    createTask.TaskLength = &temp

    body, err := proto.Marshal(createTask)

    if err != nil {
        log.Fatal("error when marshaling proto")
    }
    println("making request")
    url := "http://localhost:8080/create-task"
    bodyBuf := bytes.NewBuffer(body)
    // res, err := http.NewRequest("POST", "http://localhost:8080/create-task", bytes.NewBuffer(body))

    http.Post(url,"application/proto", bodyBuf)
    //getTaskRequest := pb.PostGetTask{
    //    TaskName: "test-task",
    //}
    // make an http request 

}

