package main

import (
	"fmt"
	_ "io"
	"log"
	"net/http"
	db "todowheel-backend/database"
	util "todowheel-backend/http-methods"
)

/*
Starts HTTP server on specified port
*/
func startServer() {

	centralDb, err := db.NewSqliteStore()
	if err != nil {
		println("unable to start db")
		log.Fatal(err)
	}
	defer centralDb.Connection.Close()

	http.HandleFunc("/health", util.GetHealth)
	http.HandleFunc("/get-task", func(w http.ResponseWriter, r *http.Request) {
		util.GetTask(w, r, centralDb)
	})
	http.HandleFunc("/get-tasks", func(w http.ResponseWriter, r *http.Request) {
		util.GetTasks(w, r, centralDb)
	})
	http.HandleFunc("/spin", func(w http.ResponseWriter, r *http.Request) {
		util.SpinTheWheel(w, r, centralDb)
	})
	http.HandleFunc("/create-task", func(w http.ResponseWriter, r *http.Request) {
		util.CreateTask(w, r, centralDb)
	})

	//http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	//    util.GetHealth(w, r)
	//})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		centralDb.Connection.Close()
		log.Fatal(err)
	}
}
