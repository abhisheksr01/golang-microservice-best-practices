package main

import (
	"github.com/abhisheksr01/golang-microservice-best-practices/api/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greetings/greet/{username}", handler.HandleGreeting)
	log.Fatal(http.ListenAndServe(":5000", r))
}
