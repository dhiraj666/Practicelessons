package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	n, err := fmt.Println("Hello Golang")
	if err != nil {
		log.Printf("error", err)
	}
	fmt.Println(n, err)
}
