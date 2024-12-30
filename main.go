package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Server")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	mux.HandleFunc("GET /history", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		fmt.Fprintf(w, "get history data by query param id: %s", id)
	})

	mux.HandleFunc("GET /history/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "get history data by id path %s", id)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
