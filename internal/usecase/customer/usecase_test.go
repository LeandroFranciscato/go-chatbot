package customer

import (
	"context"
	"reflect"
	"testing"

	"github.com/LeandroFranciscato/go-chatbot/internal/domain/entity"
	mocks "github.com/LeandroFranciscato/go-chatbot/mocks/repo"

	"go.mongodb.org/mongo-driver/bson"
)

func TestUseCase_Login(t *testing.T) {
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	type fields struct {
		repo *mocks.Repo[entity.Customer]
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Customer
		wantErr bool
		mockFn  func(repo *mocks.Repo[entity.Customer], args args)
	}{
		{
			name: "Test Login",
			fields: fields{
				repo: new(mocks.Repo[entity.Customer]),
			},
			args: args{
				ctx:      context.Background(),
				email:    "test@example.com",
				password: "password",
			},
			want:    entity.Customer{},
			wantErr: false,
			mockFn: func(repo *mocks.Repo[entity.Customer], args args) {
				repo.EXPECT().FindOne(args.ctx,
					bson.D{
						{Key: "$and", Value: bson.A{
							bson.D{{Key: "email", Value: args.email}},
							bson.D{{Key: "password", Value: args.password}}},
						},
					},
				).Return(entity.Customer{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.fields.repo, tt.args)

			usecase := useCase{
				repo: tt.fields.repo,
			}
			got, err := usecase.Login(tt.args.ctx, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("useCase.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
