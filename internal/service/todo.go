package service

import (
	"encoding/json"
	"errors"
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

func GetTodos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(todos)
}

func AddTodo(response http.ResponseWriter, request *http.Request) {
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

func GetTodo(response http.ResponseWriter, request *http.Request) {
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

func ToggleTodoStatus(response http.ResponseWriter, request *http.Request) {
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
