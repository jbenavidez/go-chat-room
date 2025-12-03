package dbrepo

import (
	pb "broker/proto/generated"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllChatMessages() ([]*pb.ChatMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select
			id, content
		from
			chat_messages
	`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var chatMessages []*pb.ChatMessage

	for rows.Next() {
		var chatMessage pb.ChatMessage
		err := rows.Scan(
			&chatMessage.Id,
			&chatMessage.Content,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		chatMessages = append(chatMessages, &chatMessage)
	}
	return chatMessages, nil
}
