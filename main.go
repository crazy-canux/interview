package main

import (
	"interview/tesla"
	"net/http"
)

func Tesla() {
	http.HandleFunc("/message", tesla.MessageHandler)
	err := http.ListenAndServe(":8081", nil)
	panic(err)
}

func main() {
	Tesla()
}
