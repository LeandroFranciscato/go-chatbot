package repo

import (
	"context"
	"reflect"
	"review-chatbot/internal/domain/entity"
	mocks "review-chatbot/mocks/repo"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_repo_Find(t *testing.T) {

	type fields struct {
		collection *mocks.Collection
	}
	type args struct {
		ctx    context.Context
		filter bson.D
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Order
		wantErr bool
		mockFn  func(collection *mocks.Collection, args args)
	}{
		{
			name: "success",
			fields: fields{
				collection: &mocks.Collection{},
			},
			args: args{
				ctx:    context.Background(),
				filter: bson.D{{Key: "key", Value: "value"}},
			},
			want:    []entity.Order{},
			wantErr: false,
			mockFn: func(collection *mocks.Collection, args args) {

				var entities []any
				cursor, err := mongo.NewCursorFromDocuments(entities, nil, nil)
				collection.EXPECT().Find(args.ctx, args.filter).Return(cursor, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.fields.collection, tt.args)

			r := &repo[entity.Order]{
				collection: tt.fields.collection,
			}
			got, err := r.Find(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_FindOne(t *testing.T) {

	type fields struct {
		collection *mocks.Collection
	}
	type args struct {
		ctx    context.Context
		filter bson.D
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Order
		wantErr bool
		mockFn  func(collection *mocks.Collection, args args)
	}{
		{
			name: "success",
			fields: fields{
				collection: &mocks.Collection{},
			},
			args: args{
				ctx:    context.Background(),
				filter: bson.D{{Key: "key", Value: "value"}},
			},
			want:    entity.Order{},
			wantErr: false,
			mockFn: func(collection *mocks.Collection, args args) {
				var entities any = entity.Order{}
				singleRes := mongo.NewSingleResultFromDocument(entities, nil, nil)
				collection.EXPECT().FindOne(args.ctx, args.filter).Return(singleRes)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.fields.collection, tt.args)

			r := &repo[entity.Order]{
				collection: tt.fields.collection,
			}
			got, err := r.FindOne(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_InsertOne(t *testing.T) {
	type fields struct {
		collection *mocks.Collection
	}
	type args struct {
		ctx    context.Context
		entity entity.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    primitive.ObjectID
		wantErr bool
		mockFn  func(collection *mocks.Collection, args args)
	}{
		{
			name: "success",
			fields: fields{
				collection: &mocks.Collection{},
			},
			args: args{
				ctx:    context.Background(),
				entity: entity.Order{},
			},
			want:    primitive.ObjectID{},
			wantErr: false,
			mockFn: func(collection *mocks.Collection, args args) {
				var a any = primitive.ObjectID{}
				collection.EXPECT().InsertOne(args.ctx, args.entity).Return(&mongo.InsertOneResult{InsertedID: a}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.fields.collection, tt.args)

			r := &repo[entity.Order]{
				collection: tt.fields.collection,
			}
			got, err := r.InsertOne(tt.args.ctx, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.InsertOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.InsertOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_InsertMany(t *testing.T) {
	type fields struct {
		collection *mocks.Collection
	}
	type args struct {
		ctx      context.Context
		entities []entity.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mockFn  func(collection *mocks.Collection, args args)
	}{
		{
			name: "success",
			fields: fields{
				collection: &mocks.Collection{},
			},
			args: args{
				ctx:      context.Background(),
				entities: []entity.Order{},
			},
			wantErr: false,
			mockFn: func(collection *mocks.Collection, args args) {
				iEntities := []any{}
				for _, entity := range args.entities {
					iEntities = append(iEntities, entity)
				}
				collection.EXPECT().InsertMany(args.ctx, iEntities).Return(nil, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.fields.collection, tt.args)

			r := &repo[entity.Order]{
				collection: tt.fields.collection,
			}
			err := r.InsertMany(tt.args.ctx, tt.args.entities)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.InsertMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_repo_DeleteMany(t *testing.T) {
	type fields struct {
		collection *mocks.Collection
	}
	type args struct {
		ctx    context.Context
		filter bson.D
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mockFn  func(collection *mocks.Collection, args args)
	}{
		{
			name: "success",
			fields: fields{
				collection: &mocks.Collection{},
			},
			args: args{
				ctx:    context.Background(),
				filter: bson.D{{Key: "key", Value: "value"}},
			},
			wantErr: false,
			mockFn: func(collection *mocks.Collection, args args) {
				collection.EXPECT().DeleteMany(args.ctx, args.filter).Return(&mongo.DeleteResult{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.fields.collection, tt.args)

			r := &repo[entity.Entity]{
				collection: tt.fields.collection,
			}
			err := r.DeleteMany(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.DeleteMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_repo_UpdateOne(t *testing.T) {
	type fields struct {
		collection *mocks.Collection
	}
	type args struct {
		ctx    context.Context
		entity entity.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mockFn  func(collection *mocks.Collection, args args)
	}{
		{
			name: "success",
			fields: fields{
				collection: &mocks.Collection{},
			},
			args: args{
				ctx:    context.Background(),
				entity: entity.Order{},
			},
			wantErr: false,
			mockFn: func(collection *mocks.Collection, args args) {
				collection.EXPECT().UpdateOne(args.ctx,
					bson.D{
						{Key: "_id", Value: args.entity.GetID()},
					},
					bson.D{
						{Key: "$set", Value: args.entity},
					},
				).Return(&mongo.UpdateResult{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.fields.collection, tt.args)

			r := &repo[entity.Order]{
				collection: tt.fields.collection,
			}
			err := r.UpdateOne(tt.args.ctx, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.UpdateOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
