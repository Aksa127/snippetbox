package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

/*
Define an application struct to hold the application-wide dependencies for the
web application. For now we'll only include the structured logger, but we'll
add more to this as the build progresses.
*/
type application struct {
	logger *slog.Logger
}

func main() {
	/*
	   Define a new command-line flag with the name 'addr', a default value of ":4000"
	   and some short help text explaining what the flag controls. The value of the
	   flag will be stored in the addr variable at runtime.
	*/
	addr := flag.String("addr", ":4000", "HTTP network address")

	dsn := flag.String("dsn", "host=localhost port=5432 user=web password=admin dbname=snippetbox sslmode=disable", "data source")

	/*
	   Importantly, we use the flag.Parse() function to parse the command-line flag.
	   This reads in the command-line flag value and assigns it to the addr
	   variable. You need to call this *before* you use the addr variable
	   otherwise it will always contain the default value of ":4000". If any errors are
	   encountered during parsing the application will be terminated.
	*/
	flag.Parse()

	/*
	   Use the slog.New() function to initialize a new structured logger, which
	   writes to the standard out stream and uses the default settings.
	*/
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	/*
		To keep the main() function tidy I've put the code for creating a connection
		pool into the separate openDB() function below. We pass openDB() the DSN
		from the command-line flag.
	*/
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	/*
	   We also defer a call to db.Close(), so that the connection pool is closed
	   before the main() function exits.
	*/
	defer db.Close()

	/*
	   Initialize a new instance of our application struct, containing the
	   dependencies (for now, just the structured logger).
	*/
	app := &application{
		logger: logger,
	}

	/*
	   The value returned from the flag.String() function is a pointer to the flag
	   value, not the value itself. So in this code, that means the addr variable
	   is actually a pointer, and we need to dereference it (i.e. prefix it with
	   the * symbol) before using it. Note that we're using the log.Printf()
	   function to interpolate the address with the log message.
	*/
	// log.Printf("starting server on %s", *addr)

	/*
	   Use the Info() method to log the starting server message at Info severity
	   (along with the listen address as an attribute).
	*/
	logger.Info("starting server", "addr", *addr)

	/*
		Use the http.ListenAndServe() function to start a new web server. We pass in
		two parameters: the TCP network address to listen on (in this case ":4000")
		and the servemux we just created. If http.ListenAndServe() returns an error
		we use the log.Fatal() function to log the error message and exit. Note
		that any error returned by http.ListenAndServe() is always non-nil.
		Call the new app.routes() method to get the servemux containing our routes,
		and pass that to http.ListenAndServe().
	*/

	/*
	   Because the err variable is now already declared in the code above, we need
	   to use the assignment operator = here, instead of the := 'declare and assign'
	   operator.
	*/
	err = http.ListenAndServe(*addr, app.routes())

	/*
	   And we also use the Error() method to log any error message returned by
	   http.ListenAndServe() at Error severity (with no additional attributes),
	   and then call os.Exit(1) to terminate the application with exit code 1.
	*/
	logger.Error(err.Error())
	os.Exit(1)

	// log.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgresql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
