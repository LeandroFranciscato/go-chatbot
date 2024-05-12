package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entity interface {
	GetID() primitive.ObjectID
	GetCollectionName() string
}
