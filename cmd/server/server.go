package server

import (
	"context"
	"fmt"
	"github.com/SunilAtAnzx/todoInGo/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func Run() {

	fmt.Println("Staring Server")
	ctx, stopHandler := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopHandler()

	router := mux.NewRouter()

	router.HandleFunc("/todos", service.GetTodos).Methods("GET")
	router.HandleFunc("/todos", service.AddTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", service.GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", service.ToggleTodoStatus).Methods("PATCH")
	router.HandleFunc("/todos/shutdown", service.ToggleTodoStatus).Methods("PATCH")

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s", ":8181"),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server terminated unexpectedly.", err)
		}
	}()

	<-ctx.Done()

	stopHandler()
	ctx, cancelHandler := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelHandler()

	err := srv.Shutdown(ctx)
	if err != nil {
		fmt.Println("Server forced to shut down..", err)
	}
	fmt.Println("Server exiting >>", ctx)
}
