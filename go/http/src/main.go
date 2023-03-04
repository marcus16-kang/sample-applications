package main

import (
    "fmt"

    "net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprint(writer, "Hello, World!")
}

func main() {
    http.HandleFunc("", index)
	
    http.ListenAndServe(":8080", nil)
}