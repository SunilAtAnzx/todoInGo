package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/SunilAtAnzx/todoInGo/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func Run() {

	var port int
	flag.IntVar(&port, "p", 8181, "Provide a port number")
	flag.Parse()

	ctx, stopHandler := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopHandler()

	router := mux.NewRouter()

	router.HandleFunc("/todos", service.GetTodos).Methods("GET")
	router.HandleFunc("/todos", service.AddTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", service.GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", service.ToggleTodoStatus).Methods("PATCH")
	router.HandleFunc("/todos/shutdown", service.ToggleTodoStatus).Methods("PATCH")

	srv := &http.Server{
		Addr:    fmt.Sprint(":", port),
		Handler: router,
	}

	fmt.Println("Server started in port :", srv.Addr)

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
