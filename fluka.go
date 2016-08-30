package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World !")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
