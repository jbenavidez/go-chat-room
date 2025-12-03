package main

import (
	pb "backend/proto/generated"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = 8080

type application struct {
	GRPCClient pb.ChatMessagesServiceClient
}

func main() {
	fmt.Println("starting back end.......")
	var app application
	// listten to websocket channel
	go ListenToWsChannel()
	//set gRPC connection
	conn, err := grpc.Dial("broker-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	//set gRPC Client
	client := pb.NewChatMessagesServiceClient(conn)
	app.GRPCClient = client

	log.Println("Starting back-end on port ", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
