package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	port = flag.Int("port", 8088, "port to listen on")
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main() {

	fmt.Println("Starting server")

	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fprintf, err := fmt.Fprintf(w, "Welcome to the home page!")
		if err != nil {
			return
		}
		fmt.Println(fprintf)
	})

	server := &http.Server{
		Addr:        ":" + strconv.Itoa(*port),
		Handler:     mux,
		ReadTimeout: 5 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
