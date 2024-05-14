package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusReturned  OrderStatus = "returned"
)

type Order struct {
	ID       primitive.ObjectID `json:"id"        bson:"_id,omitempty"`
	Customer Customer           `json:"customer"`
	Items    []Item             `json:"items"`
	Status   OrderStatus        `json:"status"`
}

func (order Order) GetID() primitive.ObjectID {
	return order.ID
}

func (order Order) GetCollectionName() string {
	return "order"
}
