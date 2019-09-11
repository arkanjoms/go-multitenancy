package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"go-multitenancy/controller"
	"log"
	"net/http"
	"os"
)

var (
	todoController controller.TodoController
	port           string
	pathPrefix     string
)

func init() {
	_ = gotenv.Load()
	port = os.Getenv("PORT")
	pathPrefix = os.Getenv("PATH_PREFIX")
	todoController = controller.TodoController{}
}

func main() {
	if port == "" {
		port = "8080"
	}

	if pathPrefix == "" {
		pathPrefix = "/go-multitenancy/{tenant}"
	}

	addr := fmt.Sprint(":", port)

	router := mux.NewRouter().PathPrefix(pathPrefix).Subrouter()

	router.HandleFunc("/todo", todoController.AddTodo).Methods(http.MethodPost)
	router.HandleFunc("/todo", todoController.GetAll).Methods(http.MethodGet)

	fmt.Println("Server listening on port: ", port)
	log.Fatal(http.ListenAndServe(addr, router))
}
