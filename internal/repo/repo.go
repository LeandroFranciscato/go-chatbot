package repo

import (
	"context"
	"errors"
	"review-chatbot/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo[T entity.Entity] struct {
	collection *mongo.Collection
}

func New[T entity.Entity](uri, user, pass, db, collection string) (repo repo[T], err error) {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().
			ApplyURI(uri).
			SetAuth(options.Credential{
				Username: user,
				Password: pass,
			}),
	)
	if err != nil {
		return repo, errors.New("error creating mongo client: " + err.Error())
	}
	repo.collection = client.Database(db).Collection(collection)
	return repo, nil
}

func (repo repo[T]) Find(ctx context.Context, filter bson.D) (records []T, err error) {
	records = []T{}
	cur, err := repo.collection.Find(ctx, filter)
	defer func() {
		_ = cur.Close(ctx)
	}()
	if err != nil {
		return records, errors.New("error finding objects:" + err.Error())
	}
	if err = cur.All(ctx, &records); err != nil {
		return records, errors.New("error parsing results: " + err.Error())
	}
	return records, nil
}
func (repo repo[T]) FindOne(ctx context.Context, filter bson.D) (result T, err error) {
	err = repo.collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return result, errors.New("error finding object:" + err.Error())
		}
	}
	return result, nil
}
