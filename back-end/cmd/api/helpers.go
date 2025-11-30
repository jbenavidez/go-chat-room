package main

import (
	"fmt"
	"log"
)

func ListenForWs(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Errror", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		} else {
			fmt.Println("Sending payload to channel", payload)
			payload.Conn = *conn
			wsChan <- payload // send payload to channel
		}
	}
}

func ListenToWsChannel() {

	for {
		e := <-wsChan // read paylod from channel
		fmt.Println("listning fo webhook event")
		switch e.Action {
		case "username":
			fmt.Println("the payload", e)
		}
	}
}
