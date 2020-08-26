package main

import (
	"interview/tesla"
	"net/http"
)

func main() {
	http.HandleFunc("/message", tesla.MessageHandler)
	err := http.ListenAndServe(":8081", nil)
	panic(err)
}
