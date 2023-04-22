package service

import (
	"bytes"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func toReader(content string) io.Reader {
	return bytes.NewBuffer([]byte(content))
}

func checkStatusCode(code int, want int, t *testing.T) {
	if code != want {
		t.Errorf("Wrong status code: got %v want %v", code, want)
	}
}

func TestGetTodos(t *testing.T) {
	r, _ := http.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTodos)
	handler.ServeHTTP(w, r)
	checkStatusCode(w.Code, http.StatusOK, t)
}

func TestGetTodosByIdForStatusNotFound(t *testing.T) {
	r, _ := http.NewRequest("GET", "/todos/id", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "10"})
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTodo)
	handler.ServeHTTP(w, r)
	checkStatusCode(w.Code, http.StatusNotFound, t)
}

func TestGetTodosByIdWhenId(t *testing.T) {
	r, _ := http.NewRequest("GET", "/todos/id", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTodo)
	handler.ServeHTTP(w, r)
	checkStatusCode(w.Code, http.StatusOK, t)
}

func TestAddTodos(t *testing.T) {
	var rqBody = toReader(`{"id": "10","item": "Clean Room1","completed": false}`)
	r, _ := http.NewRequest("POST", "/todos", rqBody)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(AddTodo)
	handler.ServeHTTP(w, r)
	checkStatusCode(w.Code, http.StatusCreated, t)
}

func TestAddTodosForStatusBadRequest(t *testing.T) {
	var rqBody = toReader(`{"id": "10","item": "Clean Room1","completed": false,"isError": true,}`)
	r, _ := http.NewRequest("POST", "/todos", rqBody)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(AddTodo)
	handler.ServeHTTP(w, r)
	checkStatusCode(w.Code, http.StatusBadRequest, t)
}

func TestToggleTodoStatus(t *testing.T) {
	r, _ := http.NewRequest("PATCH", "/todos/id", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(ToggleTodoStatus)
	handler.ServeHTTP(w, r)
	checkStatusCode(w.Code, http.StatusOK, t)
}

func TestToggleTodoStatusForStatusNotFound(t *testing.T) {
	r, _ := http.NewRequest("PATCH", "/todos/id", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "100"})
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(ToggleTodoStatus)
	handler.ServeHTTP(w, r)
	checkStatusCode(w.Code, http.StatusNotFound, t)
}
