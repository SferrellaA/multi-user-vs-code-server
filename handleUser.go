package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
)

// userList returns a list of all valid users
func userList() (users []string) {
	// Check what users folders are in the home folder
	items, err := ioutil.ReadDir(homeFolder)
	if nil != err {
		return
	}

	// Accumulate folder names that are also valid users
	for _, i := range items {
		if !i.IsDir() {
			continue
		}
		if _, _, err := getUser(i.Name()); nil == err {
			users = append(users, i.Name())
		}
	}

	return // the list of valid users with folders
}

// getUser returns a specified user's uid and gid, or an error if not found
func getUser(username string) (uid string, gid string, err error) {
	u, err := user.Lookup(username)
	if nil == err {
		uid = u.Uid
		gid = u.Gid
	}
	return
}

// getUserHash returns the password hash of the specified user
func getUserHash(username string) (hash string, err error) {
	return "AAAAAAAAAAA", nil
}

// getUserPort returns the port a specified user's environment will run on
func getUserPort(username string) (port int, err error) {
	return 8443, nil
}

// generateCommand produces the command for the user's environment
func generateCommand(username string) (command string, err error) {
	puid, pgid, err := getUser(username)
	if nil != err {
		return
	}

	hash, err := getUserHash(username)
	if nil != err {
		return
	}

	port, err := getUserPort(username)
	if nil != err {
		return
	}

	command = "docker run -d " +
		fmt.Sprintf("--name=code-server-%s ", username) +
		fmt.Sprintf("-e PUID=%s ", puid) +
		fmt.Sprintf("-e PGID=%s ", pgid) +
		fmt.Sprintf("-e TZ=Europe/London ") +
		fmt.Sprintf("-e HASHED_PASSWORD=%s ", hash) +
		fmt.Sprintf("-p %d:8443", port)

	return
}
