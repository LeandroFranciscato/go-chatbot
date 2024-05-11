package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID       primitive.ObjectID `json:"id"        bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

func (customer Customer) ToString() string {
	return customer.ID.String() + " - " + customer.Name
}
