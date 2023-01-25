package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my api")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(tasks)
}

func getoneTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Id invalido")
		return
	}

	for _, task := range tasks {
		if task.ID == taskId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}

}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks)
	router.HandleFunc("/create", createTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", router))
}
