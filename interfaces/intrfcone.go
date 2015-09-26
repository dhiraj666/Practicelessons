package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Println("sending user email to", u.name, u.email)
}

func main() {
	u := user{
		name:  "dhirajdas",
		email: "dhirajdas.666@gmail.com",
	}
	sendNotification(&u)

}

func sendNotification(n notifier) {
	n.notify()
}
