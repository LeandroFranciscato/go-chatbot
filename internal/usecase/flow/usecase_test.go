package flow

import (
	"context"
	_ "embed"
	"errors"
	"reflect"
	"review-chatbot/internal/domain/api"
	"review-chatbot/internal/domain/entity"
	"review-chatbot/internal/repo"
	"review-chatbot/internal/util"
	mocks "review-chatbot/mocks/repo"
	utilmocks "review-chatbot/mocks/util"
	"strings"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNew(t *testing.T) {
	type args struct {
		flowJson        []byte
		chatHistoryRepo repo.Repo[entity.Chat]
	}
	tests := []struct {
		name    string
		args    args
		want    Flow
		wantErr string
		mockFn  func(args)
	}{
		{
			name: "error parsing flow",
			args: args{
				flowJson: []byte("non valid json"),
			},
			want: useCase{
				timer: util.NewTimer(),
			},
			wantErr: "error parsing flow",
			mockFn:  func(a args) {},
		},
		{
			name: "error parsing stopWords",
			args: args{
				flowJson: []byte(`{"ID":"1"}`),
			},
			want: useCase{
				flow:  api.Flow{ID: "1"},
				timer: util.NewTimer(),
			},
			wantErr: "error parsing stopWords",
			mockFn: func(a args) {
				stopWordsJson = []byte("non valid json")
			},
		},
		{
			name: "success",
			args: args{
				flowJson: []byte(`{"ID":"1"}`),
			},
			want: useCase{
				flow:      api.Flow{ID: "1"},
				stopWords: map[string]struct{}{"a": {}},
				timer:     util.NewTimer(),
			},
			mockFn: func(a args) {
				stopWordsJson = []byte(`{"a": {}}`)
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			got, err := New(tt.args.flowJson, tt.args.chatHistoryRepo)
			if (err != nil) && tt.wantErr != "" && !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_ID(t *testing.T) {
	type fields struct {
		stopWords       map[string]struct{}
		flow            api.Flow
		chatHistoryRepo repo.Repo[entity.Chat]
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				flow: api.Flow{ID: "1"},
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := useCase{
				stopWords:       tt.fields.stopWords,
				flow:            tt.fields.flow,
				chatHistoryRepo: tt.fields.chatHistoryRepo,
			}
			if got := usecase.ID(); got != tt.want {
				t.Errorf("useCase.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_FinalStep(t *testing.T) {
	type fields struct {
		stopWords       map[string]struct{}
		flow            api.Flow
		chatHistoryRepo repo.Repo[entity.Chat]
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "success",
			fields: fields{
				flow: api.Flow{FinalStep: 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := useCase{
				stopWords:       tt.fields.stopWords,
				flow:            tt.fields.flow,
				chatHistoryRepo: tt.fields.chatHistoryRepo,
			}
			if got := usecase.FinalStep(); got != tt.want {
				t.Errorf("useCase.FinalStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_Name(t *testing.T) {
	type fields struct {
		stopWords       map[string]struct{}
		flow            api.Flow
		chatHistoryRepo repo.Repo[entity.Chat]
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				flow: api.Flow{Name: "flow"},
			},
			want: "flow",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := useCase{
				stopWords:       tt.fields.stopWords,
				flow:            tt.fields.flow,
				chatHistoryRepo: tt.fields.chatHistoryRepo,
			}
			if got := usecase.Name(); got != tt.want {
				t.Errorf("useCase.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_Ask(t *testing.T) {
	type fields struct {
		stopWords       map[string]struct{}
		flow            api.Flow
		chatHistoryRepo repo.Repo[entity.Chat]
	}
	type args struct {
		step int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "success",
			fields: fields{
				flow: api.Flow{
					Steps: map[string]api.Step{
						"1": {Question: "question"},
					},
				},
			},
			args: args{
				step: 1,
			},
			want: "question",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := useCase{
				stopWords:       tt.fields.stopWords,
				flow:            tt.fields.flow,
				chatHistoryRepo: tt.fields.chatHistoryRepo,
			}
			if got := usecase.Ask(tt.args.step); got != tt.want {
				t.Errorf("useCase.Ask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_Answer(t *testing.T) {
	type fields struct {
		stopWords       map[string]struct{}
		flow            api.Flow
		chatHistoryRepo repo.Repo[entity.Chat]
	}
	type args struct {
		step       int
		userAnswer string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  string
	}{
		{
			name: "success",
			fields: fields{
				flow: api.Flow{
					Steps: map[string]api.Step{
						"1": {
							IntentsDataset: map[string]string{"answer": "intent"},
							IntentConfig: map[string]api.Intent{
								"intent": {NextStep: "2", Answer: "answer"},
							},
						},
					},
				},
			},
			args: args{
				userAnswer: "answer",
				step:       1,
			},
			want:  2,
			want1: "answer",
		},
		{
			name: "final step",
			fields: fields{
				flow: api.Flow{
					FinalStep: 199,
					Steps: map[string]api.Step{
						"1": {
							IntentsDataset: map[string]string{"answer": "intent"},
							IntentConfig:   map[string]api.Intent{},
						},
					},
				},
			},
			args: args{
				userAnswer: "answer",
				step:       1,
			},
			want: 199,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := useCase{
				stopWords:       tt.fields.stopWords,
				flow:            tt.fields.flow,
				chatHistoryRepo: tt.fields.chatHistoryRepo,
			}
			got, got1 := usecase.Answer(tt.args.step, tt.args.userAnswer)
			if got != tt.want {
				t.Errorf("useCase.Answer() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("useCase.Answer() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_useCase_SaveHistory(t *testing.T) {
	type fields struct {
		stopWords       map[string]struct{}
		flow            api.Flow
		chatHistoryRepo *mocks.Repo[entity.Chat]
		timer           *utilmocks.Time
	}
	type args struct {
		ctx        context.Context
		step       int
		customerID string
		orderID    string
		history    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
		mockFn  func(mock *mocks.Repo[entity.Chat], utilmock *utilmocks.Time, args args)
	}{
		{
			name: "error finding chat history",
			fields: fields{
				chatHistoryRepo: &mocks.Repo[entity.Chat]{},
				timer:           &utilmocks.Time{},
			},
			args: args{
				ctx:        context.Background(),
				step:       1,
				customerID: "1",
				orderID:    "1",
				history:    "history",
			},
			wantErr: "error finding chat history",
			mockFn: func(mock *mocks.Repo[entity.Chat], utilmock *utilmocks.Time, args args) {
				customerObjID, _ := primitive.ObjectIDFromHex(args.customerID)
				orderObjID, _ := primitive.ObjectIDFromHex(args.orderID)

				mock.EXPECT().FindOne(args.ctx,
					bson.D{
						{Key: "$and", Value: bson.A{
							bson.D{{Key: "customer_id", Value: customerObjID}},
							bson.D{{Key: "order_id", Value: orderObjID}}},
						},
					},
				).Return(entity.Chat{}, errors.New("fake error"))
			},
		},
		{
			name: "error inserting chat history",
			fields: fields{
				chatHistoryRepo: &mocks.Repo[entity.Chat]{},
				timer:           &utilmocks.Time{},
			},
			args: args{
				ctx:        context.Background(),
				step:       1,
				customerID: "1",
				orderID:    "1",
				history:    "history",
			},
			wantErr: "error inserting chat history",
			mockFn: func(mock *mocks.Repo[entity.Chat], utilmock *utilmocks.Time, args args) {
				customerObjID, _ := primitive.ObjectIDFromHex(args.customerID)
				orderObjID, _ := primitive.ObjectIDFromHex(args.orderID)

				mock.EXPECT().FindOne(args.ctx,
					bson.D{
						{Key: "$and", Value: bson.A{
							bson.D{{Key: "customer_id", Value: customerObjID}},
							bson.D{{Key: "order_id", Value: orderObjID}}},
						},
					},
				).Return(entity.Chat{}, nil)

				rightNow := time.Now()
				utilmock.EXPECT().Now().Return(rightNow)

				mock.EXPECT().InsertOne(args.ctx, entity.Chat{
					CustomerID:  customerObjID,
					OrderID:     orderObjID,
					History:     args.history,
					Status:      entity.ChatStatusInProgress,
					Timestamp:   rightNow,
					CurrentStep: args.step,
				}).Return(primitive.ObjectID{}, errors.New("fake error"))
			},
		},
		{
			name: "success inserting chat history",
			fields: fields{
				chatHistoryRepo: &mocks.Repo[entity.Chat]{},
				timer:           &utilmocks.Time{},
			},
			args: args{
				ctx:        context.Background(),
				step:       1,
				customerID: "1",
				orderID:    "1",
				history:    "history",
			},
			wantErr: "",
			mockFn: func(mock *mocks.Repo[entity.Chat], utilmock *utilmocks.Time, args args) {
				customerObjID, _ := primitive.ObjectIDFromHex(args.customerID)
				orderObjID, _ := primitive.ObjectIDFromHex(args.orderID)

				mock.EXPECT().FindOne(args.ctx,
					bson.D{
						{Key: "$and", Value: bson.A{
							bson.D{{Key: "customer_id", Value: customerObjID}},
							bson.D{{Key: "order_id", Value: orderObjID}}},
						},
					},
				).Return(entity.Chat{}, nil)

				rightNow := time.Now()
				utilmock.EXPECT().Now().Return(rightNow)

				mock.EXPECT().InsertOne(args.ctx, entity.Chat{
					CustomerID:  customerObjID,
					OrderID:     orderObjID,
					History:     args.history,
					Status:      entity.ChatStatusInProgress,
					Timestamp:   rightNow,
					CurrentStep: args.step,
				}).Return(primitive.ObjectID{}, nil)
			},
		},
		{
			name: "error updating chat history",
			fields: fields{
				chatHistoryRepo: &mocks.Repo[entity.Chat]{},
				timer:           &utilmocks.Time{},
				flow: api.Flow{
					FinalStep: 199,
				},
			},
			args: args{
				ctx:        context.Background(),
				step:       199,
				customerID: "1",
				orderID:    "1",
				history:    "history",
			},
			wantErr: "error updating chat history",
			mockFn: func(mock *mocks.Repo[entity.Chat], utilmock *utilmocks.Time, args args) {
				customerObjID, _ := primitive.ObjectIDFromHex(args.customerID)
				orderObjID, _ := primitive.ObjectIDFromHex(args.orderID)

				chatID := primitive.NewObjectID()

				mock.EXPECT().FindOne(args.ctx,
					bson.D{
						{Key: "$and", Value: bson.A{
							bson.D{{Key: "customer_id", Value: customerObjID}},
							bson.D{{Key: "order_id", Value: orderObjID}}},
						},
					},
				).Return(entity.Chat{ID: chatID}, nil)

				mock.EXPECT().UpdateOne(args.ctx, entity.Chat{
					ID:          chatID,
					CustomerID:  customerObjID,
					OrderID:     orderObjID,
					History:     args.history,
					Status:      entity.ChatStatusDone,
					CurrentStep: args.step,
				}).Return(errors.New("fake error"))
			},
		},
		{
			name: "success updating chat history",
			fields: fields{
				chatHistoryRepo: &mocks.Repo[entity.Chat]{},
				timer:           &utilmocks.Time{},
				flow: api.Flow{
					FinalStep: 199,
				},
			},
			args: args{
				ctx:        context.Background(),
				step:       199,
				customerID: "1",
				orderID:    "1",
				history:    "history",
			},
			wantErr: "",
			mockFn: func(mock *mocks.Repo[entity.Chat], utilmock *utilmocks.Time, args args) {
				customerObjID, _ := primitive.ObjectIDFromHex(args.customerID)
				orderObjID, _ := primitive.ObjectIDFromHex(args.orderID)

				chatID := primitive.NewObjectID()

				mock.EXPECT().FindOne(args.ctx,
					bson.D{
						{Key: "$and", Value: bson.A{
							bson.D{{Key: "customer_id", Value: customerObjID}},
							bson.D{{Key: "order_id", Value: orderObjID}}},
						},
					},
				).Return(entity.Chat{ID: chatID}, nil)

				mock.EXPECT().UpdateOne(args.ctx, entity.Chat{
					ID:          chatID,
					CustomerID:  customerObjID,
					OrderID:     orderObjID,
					History:     args.history,
					Status:      entity.ChatStatusDone,
					CurrentStep: args.step,
				}).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.fields.chatHistoryRepo, tt.fields.timer, tt.args)
			usecase := useCase{
				stopWords:       tt.fields.stopWords,
				flow:            tt.fields.flow,
				chatHistoryRepo: tt.fields.chatHistoryRepo,
				timer:           tt.fields.timer,
			}
			err := usecase.SaveHistory(tt.args.ctx, tt.args.step, tt.args.customerID, tt.args.orderID, tt.args.history)
			if (err != nil) && tt.wantErr != "" && !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("useCase.SaveHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
