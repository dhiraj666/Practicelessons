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
	fmt.Println("sending username and email to ", u.name, u.email)
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Println("Sending admin username and email to ", a.name, a.email)
}

func main() {

	usr := user{
		name:  "dhiraj das",
		email: "dhirajdas.666@gmail.com",
	}
	sendNotif(&usr)
	admn := admin{
		name:  "admin dhiraj",
		email: "contact@dkdsoft.com",
	}
	sendNotif(&admn)

}

func sendNotif(n notifier) {
	n.notify()
}
