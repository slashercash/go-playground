package main

import "fmt"

type User struct {
	Name         string
	FailedLogins int
	Locked       bool
}

func NewUser(name string) User {
	return User{
		Name: name,
	}
}

func Fail(user *User) {
	user.FailedLogins += 1
	if user.FailedLogins > 0 {
		user.Locked = true
	}
}

func Reset(user *User) {
	user.FailedLogins = 0
	user.Locked = false
}

func (user *User) Reset() {
	user.FailedLogins = 0
	user.Locked = false
}

func main() {
	myUser := NewUser("Hans")
	fmt.Println("Initial User:          ", myUser)

	Fail(&myUser)
	fmt.Println("Fail User:             ", myUser)

	Reset(&myUser)
	fmt.Println("Reset User by function:", myUser)

	Fail(&myUser)
	fmt.Println("Fail User:             ", myUser)

	myUser.Reset()
	fmt.Println("Reset User by method:  ", myUser)
}
