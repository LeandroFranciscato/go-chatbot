package chat

import (
	"context"
	"review-chatbot/internal/domain/entity"
	"review-chatbot/internal/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Chat interface {
	FindByCustomer(ctx context.Context, customerID primitive.ObjectID) ([]entity.Chat, error)
}

type useCase struct {
	repo repo.Repo[entity.Chat]
}

func New(repo repo.Repo[entity.Chat]) Chat {
	return &useCase{
		repo: repo,
	}
}

func (u *useCase) FindByCustomer(ctx context.Context, customerID primitive.ObjectID) ([]entity.Chat, error) {
	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}})
	return u.repo.Find(ctx,
		bson.D{
			{Key: "customer_id", Value: customerID},
		},
		opts,
	)

}
