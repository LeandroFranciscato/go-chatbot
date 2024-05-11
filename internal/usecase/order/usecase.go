package order

import (
	"context"
	"review-chatbot/internal/domain/entity"
	"review-chatbot/internal/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order interface {
	FindByCustomer(ctx context.Context, customer primitive.ObjectID) ([]entity.Order, error)
	UpdateOne(ctx context.Context, order entity.Order) error
	FindOne(ctx context.Context, customerID primitive.ObjectID, orderID primitive.ObjectID) (entity.Order, error)
}

type useCase struct {
	repo repo.Repo[entity.Order]
}

func New(repo repo.Repo[entity.Order]) Order {
	return useCase{repo}
}

func (usecase useCase) FindByCustomer(ctx context.Context, customerID primitive.ObjectID) ([]entity.Order, error) {
	return usecase.repo.Find(ctx,
		bson.D{
			{Key: "customer._id", Value: customerID},
		},
	)
}

func (usecase useCase) UpdateOne(ctx context.Context, order entity.Order) error {
	return usecase.repo.UpdateOne(ctx, order)
}

func (usecase useCase) FindOne(ctx context.Context, customerID primitive.ObjectID, orderID primitive.ObjectID) (entity.Order, error) {
	return usecase.repo.FindOne(ctx,
		bson.D{
			{Key: "$and", Value: bson.A{
				bson.D{{Key: "customer._id", Value: customerID}},
				bson.D{{Key: "_id", Value: orderID}}},
			},
		},
	)
}
