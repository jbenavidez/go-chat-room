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
	var response WsJsonResponse

	for {
		e := <-wsChan // read paylod from channel
		fmt.Println("listning fo webhook event")
		switch e.Action {

		case "username":
			fmt.Println("the payload", e)
			//add user to online list
			clients[e.Conn] = e.Username
			users := GetOnlineusers()
			//update user list on front-end
			response.Action = "online_users"
			response.ConnectedUser = users
			broadcastToAllConn(response)

		case "chat_message":
			response.Action = "chat_message"
			// set message
			response.Message = fmt.Sprintf("<b>%s</b>: %s", e.Username, e.Message)
			broadcastToAllConn(response)
		}
	}
}

func GetOnlineusers() []string {
	var users []string
	for _, user := range clients {
		users = append(users, user)
	}
	return users
}

func broadcastToAllConn(response WsJsonResponse) {

	for client := range clients {
		err := client.WriteJSON(response)

		if err != nil {
			log.Println("WS eerr", err)
			_ = client.Close()
			delete(clients, client) // remove  from who is active tab
		}
	}
}
