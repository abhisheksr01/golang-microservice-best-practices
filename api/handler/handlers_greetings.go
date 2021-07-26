package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"unicode"

	"github.com/abhisheksr01/golang-microservice-best-practices/api/service"
	"github.com/gorilla/mux"
)

func HandleGreeting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !isValidName(username) {
		log.Fatalln("invalid username :", username)
	}
	if !ok {
		log.Fatalln("username is missing in path parameters")
	}
	fmt.Fprintf(w, service.Greet(time.Now(), username))
}

func isValidName(name string) bool {
	for _, r := range name {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
