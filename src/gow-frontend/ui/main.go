package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("../public"))
	http.Handle("/", fileServer)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
