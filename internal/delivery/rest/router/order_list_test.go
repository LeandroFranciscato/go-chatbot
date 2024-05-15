package router

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/LeandroFranciscato/go-chatbot/internal/delivery/rest"
	"github.com/LeandroFranciscato/go-chatbot/internal/domain/entity"
	flowmocks "github.com/LeandroFranciscato/go-chatbot/mocks/usecase/flow"
	ordermocks "github.com/LeandroFranciscato/go-chatbot/mocks/usecase/order"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_router_order_list(t *testing.T) {
	type fields struct {
		router router
	}
	tests := []struct {
		name   string
		fields fields
		status int
		mockFn func(orderMock *ordermocks.Order, flowMock *flowmocks.Flow)
	}{
		{
			name: "Success",
			fields: fields{
				router{
					Server: rest.Server{
						Order:      &ordermocks.Order{},
						ReviewFlow: &flowmocks.Flow{},
					},
				},
			},
			status: http.StatusOK,
			mockFn: func(ordermock *ordermocks.Order, flowmock *flowmocks.Flow) {
				objID, _ := primitive.ObjectIDFromHex("000000000000000000000000")
				ordermock.EXPECT().FindByCustomer(mock.Anything, objID).Return([]entity.Order{{ID: objID}}, nil)
				flowmock.EXPECT().GetHistory(mock.Anything, "000000000000000000000000", "000000000000000000000000").Return(entity.Chat{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// run mock
			tt.mockFn(tt.fields.router.Server.Order.(*ordermocks.Order), tt.fields.router.Server.ReviewFlow.(*flowmocks.Flow))

			// prepare router
			tt.fields.router.Engine = gin.Default()

			// create temp html file
			fileName := "order_list.html"
			_, _ = os.Create(fileName)
			tt.fields.router.Engine.LoadHTMLFiles(fileName)
			defer func() {
				_ = os.Remove(fileName)
			}()

			// add routes
			tt.fields.router.Engine.Handle(http.MethodGet, "/portal/order/list", func(ctx *gin.Context) {
				tt.fields.router.orderListHandler(ctx, "000000000000000000000000")
			})

			// prepare request
			req, _ := http.NewRequest(http.MethodGet, "/portal/order/list", nil)

			// serve request
			w := httptest.NewRecorder()
			tt.fields.router.Engine.ServeHTTP(w, req)

			if w.Code != tt.status {
				t.Fatalf("Expected to get status %v but instead got %v\n", tt.status, w.Code)
			}
		})
	}
}
