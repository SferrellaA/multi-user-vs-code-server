package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// https://hub.docker.com/r/linuxserver/code-server
var homeFolder = "/Users"

// handleMain handles /, a list of valid user links
func handleMain(w http.ResponseWriter, r *http.Request) {
	html := "The main page<br>"

	for _, user := range userList() {
		html += fmt.Sprintf("<a href=\"/u/%s\">%s</a>", user, user)
	}

	fmt.Fprintf(w, html)
}

// handleAdmin will handle user account management
func handleAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "the admin page")
}

// handleUser handles a single user's vs code instance
func handleUser(w http.ResponseWriter, r *http.Request) {
	// Parse out the user from the path
	user := strings.TrimPrefix(r.URL.Path, "/u/")

	// Check that the user actually exists
	if _, _, err := getUser(user); nil != err {
		http.Redirect(w, r, "/", 301)
	}

	// Generate the command for the user's environment
	command, err := generateCommand(user)
	if nil != err {

	}

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
