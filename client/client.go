package client

import (
	"bytes"
	"io"
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

    res, err := http.Post(url,"application/proto", bodyBuf)
    if err != nil {
        println("error making http request")
    }
    resBody, err := io.ReadAll(res.Body)
    res.Body.Close()

    if err != nil {
        println("error reading response body")
    }
    println("got response")
    println(resBody)
    //getTaskRequest := pb.PostGetTask{
    //    TaskName: "test-task",
    //}
    // make an http request 

}

