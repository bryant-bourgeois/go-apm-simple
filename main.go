package main

import (
	"errors"
	"fmt"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	tracer.Start()
	defer tracer.Stop()
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8000"
		fmt.Println("No user port supplied, using default port 8000")
	}
	fmt.Printf("Starting server on port %s...\n", portString)
	mux := httptrace.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(fmt.Sprintf(":%s", portString), mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
