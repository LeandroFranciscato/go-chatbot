package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatHistory struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CustomerID primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	OrderID    primitive.ObjectID `json:"order_id" bson:"order_id"`
	Timestamp  time.Time          `json:"timestamp"`
	History    string             `json:"history"`
}

func (chatHistory ChatHistory) GetID() primitive.ObjectID {
	return chatHistory.ID
}

func (chatHistory ChatHistory) GetCollectionName() string {
	return "chat_history"
}
