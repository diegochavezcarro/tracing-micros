package main

import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Micro3")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/call", viewHandler)
	err := http.ListenAndServe(":8082", nil)
	log.Fatal(err)
}
