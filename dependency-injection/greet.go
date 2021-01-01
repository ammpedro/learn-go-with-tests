package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet prints a greeting for someone
func Greet(writer io.Writer, name string) {
	// Originally used Printf(format string, a ...interface{})
	// Printf just calls Fprintf(os.Stdout, format, a...)
	// Printf is not testable directly because it passes straight to os.Stdout
	// and because we cannot control what happens in os.Stdout
	// Error: cannot use os.Stdout (type *os.File) as type *bytes.Buffer in argument to Greet
	// Fprintf is used instead, since it allows use to pass an io.Writer:
	// Fprintf(w io.Writer, format string, a ...interface{})
	// both os.Stdout and bytes.Buffer both implement io.Writer
	// This implementation allows reuse of Greet in webapps
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler greets
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
