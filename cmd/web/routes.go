package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	/*
	   Use the http.NewServeMux() function to initialize a new servemux, then
	   register the home function as the handler for the "/" URL pattern.
	*/
	mux := http.NewServeMux()

	/*
	   Create a file server which serves files out of the "./ui/static" directory.
	   Note that the path given to the http.Dir function is relative to the project
	   directory root.
	*/
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	/*
	   Use the mux.Handle() function to register the file server as the handler for
	   all URL paths that start with "/static/". For matching paths, we strip the
	   "/static" prefix before the request reaches the file server.
	*/
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	/*
		Prefix the route patterns with the required HTTP method (for now, we will
		restrict all three routes to acting on GET requests).
		Swap the route declarations to use the application struct's methods as the
		handler functions.
	*/
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	// Create the new route, which is restricted to POST requests only.
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}
