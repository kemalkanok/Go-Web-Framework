package main

import (

	"net/http"
	"fmt"

)
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    //running the server

    http.ListenAndServe(":8080", nil)
}
