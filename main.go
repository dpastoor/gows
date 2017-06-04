/*
gss is a very simple static file server in go
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
	--n      : don't launch the browser automatically
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	flag "github.com/ogier/pflag"

	"github.com/skratchdot/open-golang/open"
)

// VERSION is the version
const VERSION = "1.0.0"

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	noLaunch := flag.Bool("n", false, "launch the served directory in a browser")
	flag.Parse()

	dirpath, err := filepath.Abs(*directory)

	if err != nil {
		log.Fatalf("could not find directory at %s", *directory)
	}

	http.Handle("/", http.FileServer(http.Dir(dirpath)))
	launch := !*noLaunch
	if !*noLaunch {
		go func() {
			open.Run(fmt.Sprintf("http://localhost:%s", *port))
		}()
	}

	log.Printf("gss version: %s", VERSION)
	log.Printf("Serving from directory: \n %s \n on HTTP port: %s\n", dirpath, *port)
	if launch {
		log.Printf("launching browser...")
	}
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
