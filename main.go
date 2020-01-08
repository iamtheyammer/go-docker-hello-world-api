package main

import (
	"fmt"
	"net/http"
	"os"
)

var Port, _ = os.LookupEnv("PORT")

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Golang!")
	})

	if len(Port) < 1 {
		Port = "8002"
	}

	fmt.Printf("Golang Hello World API Listening on port %s\n", Port)
	http.ListenAndServe(fmt.Sprintf(":%s", Port), nil)
}