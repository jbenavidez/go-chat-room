package main

import (
	"broker/repository"
	dbrepo "broker/repository/db_repo"
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

const (
	gRpcPort = "50001"
)

type Config struct {
	DSN string
	DB  repository.DatabaseRepo
	RDB *redis.Client
}

func main() {
	fmt.Println("starting  broker service...")
	app := Config{}
	conn := app.connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	//set up db
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	// set up Redic con
	fmt.Println("connecting to redis..........")
	rdb, err := initRedis()
	if err != nil {
		panic(err)
	}
	app.RDB = rdb
	//set up helper
	NewGrpcHelper(&app)
	// Set up gRPC
	app.gRPCListenAndServe()

}

func initRedis() (*redis.Client, error) {
	fmt.Println("connecting to redis..........")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // No password by default, set if applicable
		DB:       0,  // Default DB
	})
	ctx := context.Background()
	// Ping Redis to check connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Could not connect to Redis: %v\n", err)
		return nil, err
	}
	fmt.Printf("Connected to Redis: %s\n", pong)
	return rdb, nil
}
