package auth

import "fmt"

func LoginWithCredentials(username, password string) {
	fmt.Println("Logging in with credentials...")
	fmt.Println("Name is:", username, "\nPassword:", password)
}
