package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

/*
Define a home handler function which writes a byte slice containing
"Hello from Snippetbox" as the response body.
*/

/*
Change the signature of the home handler so it is defined as a method against
application.
*/

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	/*
	   Use the Header().Add() method to add a 'Server: Go' header to the
	   response header map. The first parameter is the header name, and
	   the second parameter is the header value.
	*/
	w.Header().Add("Server", "Go")
	// w.Write([]byte("Hello from Snippetbox"))

	/*
	   Initialize a slice containing the paths to the two files. It's important
	   to note that the file containing our base template must be the *first*
	   file in the slice.
	*/
	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
	}

	/*
		Use the template.ParseFiles() function to read the files and store the
		templates in a template set. Notice that we use ... to pass the contents
		of the files slice as variadic arguments.
	*/
	ts, err := template.ParseFiles(files...)
	if err != nil {
		/*
			Because the home handler is now a method against the application
			struct it can access its fields, including the structured logger. We'll
			use this to create a log entry at Error level containing the error
			message, also including the request method and URI as attributes to
			assist with debugging.
		*/
		// app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		// log.Print(err.Error())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, r, err) // Use the serverError() helper.
		return
	}

	/*
	   Then we use the Execute() method on the template set to write the
	   template content as the response body. The last parameter to Execute()
	   represents any dynamic data that we want to pass in.
	   Use the ExecuteTemplate() method to write the content of the "base"
	   template as the response body.
	*/

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		/*
		   And we also need to update the code here to use the structured logger
		   too.
		*/
		// app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		// log.Print(err.Error())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, r, err) // Use the serverError() helper.
	}

	/*
	   Use the ExecuteTemplate() method to write the content of the "base"
	   template as the response body.
	*/
}

/*
Change the signature of the snippetView, snippetCreate,snippetCreatePost handler so it is defined as a method
against *application.
*/
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	/*
	   Extract the value of the id wildcard from the request using r.PathValue()
	   and try to convert it to an integer using the strconv.Atoi() function. If
	   it can't be converted to an integer, or the value is less than 1, we
	   return a 404 page not found response.
	*/
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	/*
	   Use the fmt.Sprintf() function to interpolate the id value with a
	   message, then write it as the HTTP response.
	*/
	// msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	// w.Write([]byte(msg))
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Use the w.WriteHeader() method to send a 201 status code.
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
