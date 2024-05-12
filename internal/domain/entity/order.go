package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string

const (
	OrderStatusPending   Status = "pending"
	OrderStatusPaid      Status = "paid"
	OrderStatusShipped   Status = "shipped"
	OrderStatusDelivered Status = "delivered"
	OrderStatusReturned  Status = "returned"
)

type Order struct {
	ID       primitive.ObjectID `json:"id"        bson:"_id,omitempty"`
	Customer Customer           `json:"customer"`
	Items    []Item             `json:"items"`
	Status   Status             `json:"status"`
}

func (order Order) GetID() primitive.ObjectID {
	return order.ID
}

func (order Order) GetCollectionName() string {
	return "order"
}
