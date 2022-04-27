package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatalln(http.ListenAndServe(":7777", http.FileServer(http.Dir("."))))
}
