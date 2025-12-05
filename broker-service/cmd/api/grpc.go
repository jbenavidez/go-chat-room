package main

import (
	pb "broker/proto/generated"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang-jwt/jwt/v4/request"
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
	// get all messages
	allMessages, err := app.DB.AllChatMessages()
	if err != nil {
		return nil, err
	}
	return &pb.GetAllChatMessagesResponse{Result: allMessages}, nil
}

func (s *server) CreateChatMessage(ctx context.Context, request *pb.CreateChatMessageRequest) (*pb.CreateChatMessageResponse, error) {
	// get the msg from request
	theMessage := request.Payload
	// store the msg
	_, err := app.DB.CreateMessage(theMessage)
	if err != nil {
		return nil, err
	}
	//send response
	return &pb.CreateChatMessageResponse{Result: "message created"}, nil

}

func (s *server) AddUserNameToCache(context.Context, *pb.AddUserNameToCacheRequest) (*pb.AddUserNameToCacheResponse, error) {
	theUsername := request.Username
	fmt.Println("the user to cache isss", theUsername)
	return &pb.AddUserNameToCacheResponse{Result: "username added"}, nil
}
