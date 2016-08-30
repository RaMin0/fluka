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
	fmt.Println("Listening on port 3000...")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
