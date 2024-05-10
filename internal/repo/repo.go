package repo

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"review-chatbot/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo repo[T]) List(ctx context.Context) (records []T, err error) {
	records = []T{}
	cur, err := repo.collection.Find(ctx, bson.D{})
	defer func() {
		_ = cur.Close(ctx)
	}()
	if err != nil {
		panic(err)
	}
	for cur.Next(context.Background()) {
		var result bson.D
		err = cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		record := new(T)
		err = repo.parseToEntity(result, record)
		if err != nil {
			return records, err
		}
		records = append(records, *record)
	}
	return records, nil
}

func (repo repo[T]) parseToEntity(mongoBson primitive.D, entity *T) error {

	mongoBytes, _ := bson.Marshal(mongoBson)
	mapa := map[string]any{}
	err := bson.Unmarshal(mongoBytes, &mapa)
	if err != nil {
		return errors.New("error parsing bson to map:" + err.Error())
	}

	bytesMap, _ := json.Marshal(mapa)
	err = json.Unmarshal(bytesMap, entity)
	if err != nil {
		return errors.New("error parsing map to entity:" + err.Error())
	}
	return nil
}
