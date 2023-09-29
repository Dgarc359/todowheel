package client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	pb "todowheel-backend/proto"

	"google.golang.org/protobuf/proto"
)

func CreateClient() {
	var temp int32 = 20
	createTask := new(pb.PostCreateTask)
	createTask.TaskName = "task1"
	createTask.TaskId = 3
	createTask.TaskLength = &temp

	body, err := proto.Marshal(createTask)

	if err != nil {
		log.Fatal("error when marshaling proto")
	}
	println("making request")
	getTaskUrl := "http://localhost:8080/create-task"
	bodyBuf := bytes.NewBuffer(body)
	http.Post(getTaskUrl, "application/proto", bodyBuf)

	temp = 200
	createTask = new(pb.PostCreateTask)
	createTask.TaskName = "task2"
	createTask.TaskId = 3
	createTask.TaskLength = &temp

	body, err = proto.Marshal(createTask)

	if err != nil {
		log.Fatal("error when marshaling proto")
	}
	println("making request")
	getTaskUrl = "http://localhost:8080/create-task"
	bodyBuf = bytes.NewBuffer(body)
	http.Post(getTaskUrl, "application/proto", bodyBuf)

	getTasksRequest := new(pb.PostGetTasks)

	var maxTaskLength int32 = 500
	var minTaskLength int32 = 10
	getTasksRequest.MaxTaskLength = &maxTaskLength
	getTasksRequest.MinTaskLength = &minTaskLength
	body, err = proto.Marshal(getTasksRequest)
	if err != nil {
		log.Fatal("unable to marshal getTasks")
	}
	getTasksBuf := bytes.NewBuffer(body)
	getTasksUrl := "http://localhost:8080/spin"
	println("making get tasks request")
	res, err := http.Post(getTasksUrl, "application/proto", getTasksBuf)
	if err != nil {
		fmt.Println(err)
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(body))

	defer res.Body.Close()
}
