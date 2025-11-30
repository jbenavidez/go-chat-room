package main

import (
	"log"
	"net/http"
)

// wsChan channel for ws payload
var wsChan = make(chan WsPayload)

var clients = make(map[WebSocketConnection]string)

func (app *application) Tester(w http.ResponseWriter, r *http.Request) {
	resp := JSONResponse{
		Error:   false,
		Message: "Welcome",
	}
	_ = app.writeJSON(w, http.StatusOK, resp)
}

func (app *application) WsChatRoom(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("client is connected to WS")

	var response WsJsonResponse
	response.Message = `<em><small> connected to served</small></em>`
	conn := WebSocketConnection{Conn: ws}

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}
	go ListenForWs(&conn) // start go runtine to listen Ws
}
