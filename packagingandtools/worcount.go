package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dhiraj666/gocode/packagingandtools/words"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text file . \n", count)
}
