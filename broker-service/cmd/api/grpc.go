package main

import (
	pb "broker/proto/generated"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"slices"

	"github.com/go-redis/redis/v8"
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

func (s *server) AddUserNameToCache(ctx context.Context, request *pb.AddUserNameToCacheRequest) (*pb.AddUserNameToCacheResponse, error) {
	theUsername := request.Username
	// Get username-list
	var usernameList []string

	//get slide of username
	resp, err := app.RDB.Get(ctx, "users_online").Bytes()
	if errors.Is(err, redis.Nil) {
		fmt.Println("key was not found which mean no user conencted")
		//set userlist
		users_online := []string{theUsername}
		payload, err := json.Marshal(users_online)
		if err != nil {
			fmt.Println("error marshaling user_online", err)
			return nil, err
		}
		//set key
		err = app.RDB.Set(ctx, "users_online", payload, 0).Err()
		if err != nil {
			fmt.Println("error setting key")
			return nil, err
		}

		return &pb.AddUserNameToCacheResponse{Result: "username added"}, nil
	} else {
		//Unmarshal
		err = json.Unmarshal(resp, &usernameList)
		if err != nil {
			fmt.Println("error unmarshalling", err)
		}
		if slices.Contains(usernameList, theUsername) {
			fmt.Println()
		}

		//update cache
		usernameList = append(usernameList, theUsername)
		req, err := json.Marshal(usernameList)
		if err != nil {
			fmt.Println("error marshaling usernameList", err)
			return nil, err
		}
		err = app.RDB.Set(ctx, "users_online", req, 0).Err()
		if err != nil {
			fmt.Println("error setting key")
			return nil, err
		}
		fmt.Println("final_list", usernameList)
	}

	return &pb.AddUserNameToCacheResponse{Result: "username added"}, nil
}

func (s *server) GetAllConnectedusers(ctx context.Context, request *emptypb.Empty) (*pb.GetAllConnectedusersResponse, error) {
	return &pb.GetAllConnectedusersResponse{Result: []string{"heloo", "there"}}, nil
}
