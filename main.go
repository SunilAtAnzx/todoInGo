package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(todos)
}

func addTodo(response http.ResponseWriter, request *http.Request) {
	var newTodo todo
	err := json.NewDecoder(request.Body).Decode(&newTodo)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	todos = append(todos, newTodo)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(newTodo)
}

func getTodo(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	todo, err := getTodoById(id)

	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		payload := make(map[string]string)
		payload["message"] = "Todo not found"
		json.NewEncoder(response).Encode(payload)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(todo)
}

func getTodoById(id string) (*todo, error) {

	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func toggleTodoStatus(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	todo, err := getTodoById(id)

	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		payload := make(map[string]string)
		payload["message"] = "Todo not found"
		json.NewEncoder(response).Encode(payload)
		return
	}

	todo.Completed = !todo.Completed

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(todo)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos", addTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", toggleTodoStatus).Methods("PATCH")

	err := http.ListenAndServe(":5000", router)

	if err != nil {
		fmt.Println(err)
		return
	}
}
