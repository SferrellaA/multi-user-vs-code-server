package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "the main page")
}

func handleAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "the admin page")
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	// Parse out the user from the path
	user := strings.TrimPrefix(r.URL.Path, "/u/")

	// Check that the user actually exists
	if !validUser(user) {
		http.Redirect(w, r, "/", 301)
	}

	// Generate the command for the user's environment
	command := generateCommand(user)

	fmt.Fprintf(w, command)
}

func main() {
	http.HandleFunc("/admin", handleAdmin)
	http.HandleFunc("/u/", handleUser)
	http.HandleFunc("/", handleMain)
	err := http.ListenAndServe(":8081", nil)
	if nil != err {
		log.Fatal(err)
	}
}
