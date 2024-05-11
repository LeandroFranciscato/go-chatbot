package util

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"review-chatbot/internal/domain/entity"
	"review-chatbot/internal/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Migrate(ctx context.Context, uri, user, pass, db string) error {

	hash := md5.New()
	_, err := hash.Write([]byte("doe"))
	if err != nil {
		return errors.New("error hashing password: " + err.Error())
	}
	hashedPass := hex.EncodeToString(hash.Sum(nil))

	customers := []entity.Customer{
		{
			ID:       primitive.NewObjectID(),
			Name:     "John Doe",
			Email:    "john@email.com",
			Password: hashedPass,
		},
		{
			ID:       primitive.NewObjectID(),
			Name:     "Jane Doe",
			Email:    "jane@email.com",
			Password: hashedPass,
		},
	}

	Items := []entity.Item{
		{
			ID:   primitive.NewObjectID(),
			Name: "Item 1",
		},
		{
			ID:   primitive.NewObjectID(),
			Name: "Item 2",
		},
		{
			ID:   primitive.NewObjectID(),
			Name: "Item 3",
		},
	}

	Orders := []entity.Order{
		{
			ID:       primitive.NewObjectID(),
			Customer: customers[0],
			Items: []entity.Item{
				Items[0],
				Items[1],
			},
			Status: entity.OrderStatusShipped,
		},
		{
			ID:       primitive.NewObjectID(),
			Customer: customers[0],
			Items: []entity.Item{
				Items[0],
			},
			Status: entity.OrderStatusShipped,
		},
		{
			ID:       primitive.NewObjectID(),
			Customer: customers[1],
			Items: []entity.Item{
				Items[1],
				Items[2],
			},
			Status: entity.OrderStatusShipped,
		},
	}

	customerRepo, err := repo.New[entity.Customer](uri, user, pass, db, "customer")
	if err != nil {
		return errors.New("custom repo error: " + err.Error())
	}
	if err = customerRepo.DeleteMany(ctx, bson.D{}); err != nil {
		return errors.New("delete many customers error: " + err.Error())
	}
	if err = customerRepo.InsertMany(ctx, customers); err != nil {
		return errors.New("insert many customers error: " + err.Error())
	}

	itemRepo, err := repo.New[entity.Item](uri, user, pass, db, "item")
	if err != nil {
		return errors.New("item repo error: " + err.Error())
	}
	if err = itemRepo.DeleteMany(ctx, bson.D{}); err != nil {
		return errors.New("delete many items error: " + err.Error())
	}
	if err = itemRepo.InsertMany(ctx, Items); err != nil {
		return errors.New("insert many items error: " + err.Error())
	}

	orderRepo, err := repo.New[entity.Order](uri, user, pass, db, "order")
	if err != nil {
		return errors.New("order repo error: " + err.Error())
	}
	if err = orderRepo.DeleteMany(ctx, bson.D{}); err != nil {
		return errors.New("delete many orders error: " + err.Error())
	}
	if err = orderRepo.InsertMany(ctx, Orders); err != nil {
		return errors.New("insert many orders error: " + err.Error())
	}
	return nil
}
