package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var Port, _ = os.LookupEnv("PORT")

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		body := "N/A (one was not submitted or this is not a post request)"

		if r.Method == http.MethodPost {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Request recieved, but there was an error reading the body.")
				return
			}

			bs := string(b)

			if len(bs) > 0 {
				body = bs
			}
		}

		fmt.Fprintf(w, "Hello from Golang! Path: %s; Method: %s; Full URL: %s; Body: %s;", r.URL.Path, r.Method, r.URL.String(), body)
	})

	if len(Port) < 1 {
		Port = "8002"
	}

	fmt.Printf("Golang Hello World API Listening on port %s\n", Port)
	http.ListenAndServe(fmt.Sprintf(":%s", Port), nil)
}