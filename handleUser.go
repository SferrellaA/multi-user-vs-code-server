package main

import "fmt"

// validUser checks that a provided username is valid
func validUser(username string) bool {
	userList := []string{"fake"}
	for _, u := range userList {
		if username == u {
			return true
		}
	}
	return false
}

func getUserIDs(username string) (int, int) {
	return 1000, 1000
}

func getUserHash(username string) string {
	return "AAAAAAAAAAA"
}

func getUserPort(usrname string) int {
	return 8443
}

// generateCommand produces the command for the user's environment
func generateCommand(username string) string {
	puid, pgid := getUserIDs(username)
	hash := getUserHash(username)
	port := getUserPort(username)

	command := "docker run -d " +
		fmt.Sprintf("--name=code-server-%s ", username) +
		fmt.Sprintf("-e PUID=%d ", puid) +
		fmt.Sprintf("-e PGID=%d ", pgid) +
		fmt.Sprintf("-e TZ=Europe/London ") +
		fmt.Sprintf("-e HASHED_PASSWORD=%s ", hash) +
		fmt.Sprintf("-p %d:8443", port)

	return command
}
