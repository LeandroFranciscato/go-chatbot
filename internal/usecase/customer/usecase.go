package customer

import (
	"context"

	"github.com/LeandroFranciscato/go-chatbot/internal/domain/entity"
	"github.com/LeandroFranciscato/go-chatbot/internal/repo"

	"go.mongodb.org/mongo-driver/bson"
)

type Customer interface {
	Login(ctx context.Context, email, password string) (entity.Customer, error)
}

type useCase struct {
	repo repo.Repo[entity.Customer]
}

func New(repo repo.Repo[entity.Customer]) Customer {
	return useCase{repo}
}

func (usecase useCase) Login(ctx context.Context, email, password string) (entity.Customer, error) {
	return usecase.repo.FindOne(ctx,
		bson.D{
			{Key: "$and", Value: bson.A{
				bson.D{{Key: "email", Value: email}},
				bson.D{{Key: "password", Value: password}}},
			},
		},
	)
}
