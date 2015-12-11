package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Calidev!")
}

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
