package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct{}

func main() {
	fmt.Println("starting back end.......")
	var app application
	// listten to websocket channel
	go ListenToWsChannel()

	log.Println("Starting application on port ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
