package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID   primitive.ObjectID `json:"id"        bson:"_id,omitempty"`
	Name string             `json:"name"      bson:"name"`
}

func (item Item) ToString() string {
	return item.ID.String() + " - " + item.Name
}
