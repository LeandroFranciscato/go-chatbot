package repo

import (
	"context"
	"errors"
	"review-chatbot/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo[T entity.Entity] interface {
	Find(ctx context.Context, filter bson.D, opts ...*options.FindOptions) (records []T, err error)
	FindOne(ctx context.Context, filter bson.D) (T, error)
	InsertOne(ctx context.Context, entity T) (primitive.ObjectID, error)
	InsertMany(ctx context.Context, entities []T) error
	DeleteMany(ctx context.Context, filter bson.D) error
	UpdateOne(ctx context.Context, entity T) error
}

type Collection interface {
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

type repo[T entity.Entity] struct {
	collection Collection
}

func New[T entity.Entity](uri, user, pass, db, collection string) (Repo[T], error) {
	repo := repo[T]{}
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

func (repo repo[T]) Find(ctx context.Context, filter bson.D, opts ...*options.FindOptions) (records []T, err error) {
	records = []T{}
	cur, err := repo.collection.Find(ctx, filter, opts...)
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
	err = repo.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return result, errors.New("error finding object:" + err.Error())
		}
	}
	return result, nil
}

func (repo repo[T]) InsertOne(ctx context.Context, entity T) (id primitive.ObjectID, err error) {
	res, err := repo.collection.InsertOne(ctx, entity)
	if err != nil {
		return id, errors.New("error insertind entity:" + err.Error())
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (repo repo[T]) InsertMany(ctx context.Context, entities []T) error {
	iEntities := []any{}
	for _, entity := range entities {
		iEntities = append(iEntities, entity)
	}
	_, err := repo.collection.InsertMany(ctx, iEntities)
	if err != nil {
		return errors.New("error inserting entity:" + err.Error())
	}

	return nil
}

func (repo repo[T]) DeleteMany(ctx context.Context, filter bson.D) error {
	_, err := repo.collection.DeleteMany(ctx, filter)
	if err != nil {
		return errors.New("error deleting entity:" + err.Error())
	}

	return nil
}

func (repo repo[T]) UpdateOne(ctx context.Context, entity T) error {
	_, err := repo.collection.UpdateOne(
		ctx,
		bson.D{
			{Key: "_id", Value: entity.GetID()},
		},
		bson.D{
			{Key: "$set", Value: entity},
		},
	)
	if err != nil {
		return errors.New("error updating entity:" + err.Error())
	}

	return nil
}
