package chat

import (
	"context"
	"reflect"
	"testing"

	"github.com/LeandroFranciscato/go-chatbot/internal/domain/entity"
	mocks "github.com/LeandroFranciscato/go-chatbot/mocks/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestUseCase_FindByCustomer(t *testing.T) {
	type args struct {
		ctx        context.Context
		customerID primitive.ObjectID
	}
	type fields struct {
		repo *mocks.Repo[entity.Chat]
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Chat
		wantErr bool
		mockFn  func(repo *mocks.Repo[entity.Chat], args args)
	}{
		{
			name: "Test FindByCustomer",
			fields: fields{
				repo: new(mocks.Repo[entity.Chat]),
			},
			args: args{
				ctx:        context.Background(),
				customerID: primitive.NewObjectID(),
			},
			want:    []entity.Chat{},
			wantErr: false,
			mockFn: func(repo *mocks.Repo[entity.Chat], args args) {
				opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}})
				repo.EXPECT().Find(args.ctx, bson.D{
					{Key: "customer_id", Value: args.customerID},
				}, opts).Return([]entity.Chat{}, nil)
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
