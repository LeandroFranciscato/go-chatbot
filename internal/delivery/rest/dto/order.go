package dto

import "github.com/LeandroFranciscato/go-chatbot/internal/domain/entity"

type Order struct {
	entity.Order
	ChatStatus entity.ChatStatus `json:"chat_status"`
}
