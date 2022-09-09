package main

import "fmt"

type User struct {
	Name         string
	FailedLogins int
	Locked       bool
}

func main() {
	myUser := User{
		Name: "admin",
	}

	fmt.Println(myUser)

	myUser.FailedLogins += 1

	if myUser.FailedLogins > 0 {
		myUser.Locked = true
	}

	fmt.Println(myUser)
}
