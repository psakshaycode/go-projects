package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	s := "127.0.0.1:8080"
	log.Println("server active @:", s)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye world")
	})
	http.ListenAndServe(s, nil)
}
