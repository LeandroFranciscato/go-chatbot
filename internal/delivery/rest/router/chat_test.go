package router

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/LeandroFranciscato/go-chatbot/internal/delivery/rest"
	"github.com/LeandroFranciscato/go-chatbot/internal/domain/entity"
	flowmocks "github.com/LeandroFranciscato/go-chatbot/mocks/usecase/flow"
	ordermocks "github.com/LeandroFranciscato/go-chatbot/mocks/usecase/order"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func Test_router_chat(t *testing.T) {
	type fields struct {
		router router
	}
	tests := []struct {
		name   string
		fields fields
		status int
		mockFn func(flowmock *flowmocks.Flow, ordermock *ordermocks.Order)
	}{
		{
			name: "Success",
			fields: fields{
				router{
					Server: rest.Server{
						ReviewFlow: &flowmocks.Flow{},
						Order:      &ordermocks.Order{},
					},
				},
			},
			status: http.StatusOK,
			mockFn: func(flowmock *flowmocks.Flow, ordermock *ordermocks.Order) {
				flowmock.EXPECT().GetHistory(mock.Anything, "000000000000000000000000", "000000000000000000000000").Return(entity.Chat{}, nil)
				objID, _ := primitive.ObjectIDFromHex("000000000000000000000000")
				ordermock.EXPECT().FindOne(mock.Anything, objID, objID).Return(entity.Order{}, nil)
				flowmock.EXPECT().Ask(1).Return("question")
				flowmock.EXPECT().SaveHistory(mock.Anything, 1, "000000000000000000000000", "000000000000000000000000", mock.AnythingOfType("string")).Return(nil)
				flowmock.EXPECT().Name().Return("flow")
				flowmock.EXPECT().FinalStep().Return(199)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// run mock
			tt.mockFn(tt.fields.router.Server.ReviewFlow.(*flowmocks.Flow), tt.fields.router.Server.Order.(*ordermocks.Order))

			// prepare router
			tt.fields.router.Engine = gin.Default()

			// create temp html file
			fileName := "chat.html"
			_, _ = os.Create(fileName)
			tt.fields.router.Engine.LoadHTMLFiles(fileName)
			defer func() {
				_ = os.Remove(fileName)
			}()

			// add routes
			tt.fields.router.Engine.Handle(http.MethodPost, "/portal/chat/help", func(ctx *gin.Context) {
				tt.fields.router.chatHandler(ctx, tt.fields.router.Server.ReviewFlow, "000000000000000000000000")
			})

			// prepare request
			req, _ := http.NewRequest(http.MethodPost, "/portal/chat/help", nil)

			// serve request
			w := httptest.NewRecorder()
			tt.fields.router.Engine.ServeHTTP(w, req)

			if w.Code != tt.status {
				t.Fatalf("Expected to get status %v but instead got %v\n", tt.status, w.Code)
			}
		})
	}
}
