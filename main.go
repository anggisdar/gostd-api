package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Simple api running in port 8080")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		log.Println("'/status' was called")
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "OK"}`))
	})

	http.HandleFunc("/", handleRoot)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("'/' was called")
	w.WriteHeader(http.StatusNotImplemented)
}
