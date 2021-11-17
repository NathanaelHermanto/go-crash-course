package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hallo Hallo")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Starting the server...")
	http.ListenAndServe("localhost:3000", nil)

}
