package order

import (
	"context"
	"reflect"
	"review-chatbot/internal/domain/entity"
	mocks "review-chatbot/mocks/repo"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_useCase_FindByCustomer(t *testing.T) {
	type fields struct {
		repo *mocks.Repo[entity.Order]
	}
	type args struct {
		ctx        context.Context
		customerID primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Order
		wantErr bool
		mockFn  func(repo *mocks.Repo[entity.Order], args args)
	}{
		{
			name: "success",
			fields: fields{
				repo: &mocks.Repo[entity.Order]{},
			},
			args: args{
				ctx:        context.Background(),
				customerID: primitive.NewObjectID(),
			},
			want:    []entity.Order{},
			wantErr: false,
			mockFn: func(repo *mocks.Repo[entity.Order], args args) {
				repo.EXPECT().Find(context.Background(), bson.D{
					{Key: "customer._id", Value: args.customerID},
				}).Return([]entity.Order{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.fields.repo, tt.args)

			usecase := useCase{
				repo: tt.fields.repo,
			}
			got, err := usecase.FindByCustomer(tt.args.ctx, tt.args.customerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.FindByCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("useCase.FindByCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_UpdateOne(t *testing.T) {
	type args struct {
		ctx   context.Context
		order entity.Order
	}
	type fields struct {
		repo *mocks.Repo[entity.Order]
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mockFn  func(repo *mocks.Repo[entity.Order], args args)
	}{
		{
			name: "Test UpdateOne",
			fields: fields{
				repo: new(mocks.Repo[entity.Order]),
			},
			args: args{
				ctx:   context.Background(),
				order: entity.Order{},
			},
			wantErr: false,
			mockFn: func(repo *mocks.Repo[entity.Order], args args) {
				repo.EXPECT().UpdateOne(args.ctx, args.order).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.fields.repo, tt.args)

			usecase := useCase{
				repo: tt.fields.repo,
			}
			err := usecase.UpdateOne(tt.args.ctx, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.UpdateOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_FindOne(t *testing.T) {
	type args struct {
		ctx        context.Context
		customerID primitive.ObjectID
		orderID    primitive.ObjectID
	}
	type fields struct {
		repo *mocks.Repo[entity.Order]
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Order
		wantErr bool
		mockFn  func(repo *mocks.Repo[entity.Order], args args)
	}{
		{
			name: "Test FindOne",
			fields: fields{
				repo: new(mocks.Repo[entity.Order]),
			},
			args: args{
				ctx:        context.Background(),
				customerID: primitive.NewObjectID(),
				orderID:    primitive.NewObjectID(),
			},
			want:    entity.Order{},
			wantErr: false,
			mockFn: func(repo *mocks.Repo[entity.Order], args args) {
				repo.EXPECT().FindOne(args.ctx,
					bson.D{
						{Key: "$and", Value: bson.A{
							bson.D{{Key: "customer._id", Value: args.customerID}},
							bson.D{{Key: "_id", Value: args.orderID}}},
						},
					},
				).Return(entity.Order{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.fields.repo, tt.args)

			usecase := useCase{
				repo: tt.fields.repo,
			}
			got, err := usecase.FindOne(tt.args.ctx, tt.args.customerID, tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("useCase.FindOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
