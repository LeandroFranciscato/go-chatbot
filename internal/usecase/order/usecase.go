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
	Update(ctx context.Context, order entity.Order) error
}

type useCase struct {
	repo repo.Repo[entity.Order]
}

func New(repo repo.Repo[entity.Order]) Order {
	return useCase{repo}
}

func (usecase useCase) FindByCustomer(ctx context.Context, customer primitive.ObjectID) ([]entity.Order, error) {
	return usecase.repo.Find(ctx,
		bson.D{
			{Key: "customer._id", Value: customer},
		},
	)
}

func (usecase useCase) Update(ctx context.Context, order entity.Order) error {
	return usecase.repo.UpdateOne(ctx, order)
}
