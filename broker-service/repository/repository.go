package repository

import (
	pb "broker/proto/generated"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllChatMessages() ([]*pb.ChatMessage, error)
}
