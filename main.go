package main

import (
	"404Chat/controller"
	"log"
	"net/http"
)

func main() {
	r := controller.Router()

	r.HandleFunc("/login", controller.SendLogin)
	r.HandleFunc("/4chat", controller.SendClient)

	err := http.ListenAndServe(":4001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
