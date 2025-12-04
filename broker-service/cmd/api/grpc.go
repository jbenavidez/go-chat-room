package main

import (
	pb "broker/proto/generated"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

var app *Config

// NewGrpcHelper make app config available
func NewGrpcHelper(a *Config) {
	app = a
}

type server struct {
	pb.UnimplementedChatMessagesServiceServer
}

// gRPCListenAndServe set up gRPC conenction
func (app *Config) gRPCListenAndServe() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterChatMessagesServiceServer(srv, &server{})

	reflection.Register(srv)
	log.Printf("gRPC server started on port %s ", gRpcPort)
	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) GetAllChatMessages(ctx context.Context, request *emptypb.Empty) (*pb.GetAllChatMessagesResponse, error) {
	// test
	var allMessages []*pb.ChatMessage
	testMessage := pb.ChatMessage{
		Id:        1,
		Content:   "hello there",
		Username:  "Frodo",
		CreatedAt: "2025-12-04 23:32:03.786395",
	}
	allMessages = append(allMessages, &testMessage)
	return &pb.GetAllChatMessagesResponse{Result: allMessages}, nil
}
