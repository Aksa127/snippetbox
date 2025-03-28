package main

import (
	"log"
	"net/http"
)

func main() {
	/*
	   Use the http.NewServeMux() function to initialize a new servemux, then
	   register the home function as the handler for the "/" URL pattern.
	*/
	mux := http.NewServeMux()

	/*
	   Prefix the route patterns with the required HTTP method (for now, we will
	   restrict all three routes to acting on GET requests).
	*/
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// Create the new route, which is restricted to POST requests only.
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Print a log message to say that the server is starting.
	log.Print("starting server on :4000")

	/*
	   Use the http.ListenAndServe() function to start a new web server. We pass in
	   two parameters: the TCP network address to listen on (in this case ":4000")
	   and the servemux we just created. If http.ListenAndServe() returns an error
	   we use the log.Fatal() function to log the error message and exit. Note
	   that any error returned by http.ListenAndServe() is always non-nil.
	*/
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}
