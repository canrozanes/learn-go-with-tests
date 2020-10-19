package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet name
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler writes "Can" to http.Response
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Can")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
