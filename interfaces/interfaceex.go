package main

import (
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		os.Exit(-1)
	}

}

func main() {

	r, err := http.Get(os.Args[1])
	if err != nil {
		return
	}
	io.Copy(os.Stdout, r.Body)

	if err := r.Body.Close(); err != nil {
		return
	}
}
