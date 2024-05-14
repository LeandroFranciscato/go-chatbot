package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatStatus string

const (
	ChatStatusInProgress ChatStatus = "inProgress"
	ChatStatusDone       ChatStatus = "done"
)

type Chat struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CustomerID  primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	OrderID     primitive.ObjectID `json:"order_id" bson:"order_id"`
	Status      ChatStatus         `json:"status"`
	Timestamp   time.Time          `json:"timestamp"`
	CurrentStep int                `json:"current_step"`
	History     string             `json:"history"`
}

func (chatHistory Chat) GetID() primitive.ObjectID {
	return chatHistory.ID
}

func (chatHistory Chat) GetCollectionName() string {
	return "chat_history"
}
