package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"review-chatbot/internal/delivery/rest"
	"review-chatbot/internal/domain/entity"
	mocks "review-chatbot/mocks/usecase/order"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_router_order_delivered(t *testing.T) {
	type fields struct {
		router router
	}
	tests := []struct {
		name   string
		fields fields
		status int
		mockFn func(mock *mocks.Order)
	}{
		{
			name: "Success",
			fields: fields{
				router{
					Server: rest.Server{
						Order: &mocks.Order{},
					},
				},
			},
			status: http.StatusPermanentRedirect,
			mockFn: func(ordermock *mocks.Order) {
				cObjID, _ := primitive.ObjectIDFromHex("000000000000000000000000")
				oObjID, _ := primitive.ObjectIDFromHex("000000000000000000000000")
				ordermock.EXPECT().FindOne(mock.Anything, cObjID, oObjID).Return(entity.Order{}, nil)
				ordermock.EXPECT().UpdateOne(mock.Anything, entity.Order{Status: entity.OrderStatusDelivered}).Return(nil)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// run mock
			tt.mockFn(tt.fields.router.Server.Order.(*mocks.Order))

			// prepare router
			tt.fields.router.Engine = gin.Default()

			// add routes
			tt.fields.router.orderDelivered(tt.fields.router.Engine.Group("portal"))

			// prepare request
			req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/portal/customer/%s/order/%s/delivered", "000000000000000000000000", "000000000000000000000000"), nil)

			// serve request
			w := httptest.NewRecorder()
			tt.fields.router.Engine.ServeHTTP(w, req)

			if w.Code != tt.status {
				t.Fatalf("Expected to get status %v but instead got %v\n", tt.status, w.Code)
			}
		})
	}
}
