/*
Serve is a very simple static file server in go
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/skratchdot/open-golang/open"
)

// VERSION is the version
const VERSION = "0.0.2"

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))
	go func() {
		open.Run(fmt.Sprintf("http://localhost:%s", *port))
	}()

	dirpath, err := filepath.Abs(*directory)

	if err != nil {
		log.Fatalf("could not find directory at %s", *directory)
	}

	log.Printf("gss version: %s", VERSION)
	log.Printf("Serving %s \non HTTP port: %s\n", dirpath, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
