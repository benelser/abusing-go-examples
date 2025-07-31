package main

import (
	"log"
	"net/http"

	"github.com/benelser/abusing-go-examples/internal"
)

func main() {
	http.HandleFunc("/echo", internal.EchoHandler)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Println("listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
