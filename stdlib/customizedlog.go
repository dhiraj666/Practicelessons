// sample program to demonstrate on how to create customized loggers

package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // Just about anything
	Info    *log.Logger // Important information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical problem

)

func init() {
	file, err := os.OpenFile("errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about ")
	Error.Println("Something has failed")
}
