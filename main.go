package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"go-multitenancy/controller"
	"go-multitenancy/multitenancy"
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

	router.Use(multitenancy.TenantResolverMiddleware)

	router.HandleFunc("/todos", todoController.AddTodo).Methods(http.MethodPost)
	router.HandleFunc("/todos", todoController.GetAll).Methods(http.MethodGet)

	fmt.Println("Server listening on port: ", port)
	log.Fatal(http.ListenAndServe(addr, router))
}
