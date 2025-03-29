// In go, we don't need to export the function just write the fn name starting with a capital letter.
// Step 1: Run-> go mod init <package-name> (usually we give github repo link)
// To get 3rd party packages, we can use go get <package-name>
package main

import (
	"fmt"
	"github.com/karan-develops/Exploring-Go/auth"
)

func main(){
	auth.LoginWithCredentials("Karan","aggarwal")
	session := auth.SessionDetails()
	fmt.Println(session)
}