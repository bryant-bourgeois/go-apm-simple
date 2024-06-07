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

func getHostname() string {
	out, err := os.Hostname()
	if err != nil {
		return "default_hostname"
	}
	return out
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	pageResponse := fmt.Sprintf("Index route from %s\n", getHostname())
	io.WriteString(w, pageResponse)
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	pageResponse := fmt.Sprintf("Hello, from %s\n", getHostname())
	io.WriteString(w, pageResponse)
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
