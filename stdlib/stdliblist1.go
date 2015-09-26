/*
 currently there are over 100 pkgs organized within 38 categories
archive bufio bytes compress container crypto database
debug encoding errors expvar flag fmt go hash html image
index io log math mime net os path reflect regexp runtime
sort strconv strings sync syscall testing text time unicode
unsafe
http://golang.org/pkg
sourcegraph
Logging is a way to find those bugs and learn more about how our program is functioning
provides code tracing profiling and analytics
we can also custom loggers to implement your own specific logging needs
Unix architects added a device called stderr
This device was created to be the default destination for logging
separate their programs output from their logging
*/

package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// println writes to the standard logger
	log.Println("message")

	// Fatalln is println followed by a call to os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln is println folowed by a call to panic()
	log.Panicln("panic mesage")

}
