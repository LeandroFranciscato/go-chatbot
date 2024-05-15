package router

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/LeandroFranciscato/go-chatbot/internal/delivery/rest"
	"github.com/LeandroFranciscato/go-chatbot/internal/domain/entity"
	flowmocks "github.com/LeandroFranciscato/go-chatbot/mocks/usecase/flow"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func Test_router_chat_history(t *testing.T) {
	type fields struct {
		router router
	}
	tests := []struct {
		name   string
		fields fields
		status int
		mockFn func(flowmock *flowmocks.Flow)
	}{
		{
			name: "Success",
			fields: fields{
				router{
					Server: rest.Server{
						ReviewFlow: &flowmocks.Flow{},
					},
				},
			},
			status: http.StatusOK,
			mockFn: func(flowmock *flowmocks.Flow) {
				flowmock.EXPECT().GetHistory(mock.Anything, "000000000000000000000000", "000000000000000000000000").Return(entity.Chat{}, nil)
				flowmock.EXPECT().Name().Return("")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// run mock
			tt.mockFn(tt.fields.router.Server.ReviewFlow.(*flowmocks.Flow))

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
			tt.fields.router.chatHistory(tt.fields.router.Engine.Group("/portal/chat"))

			// prepare request
			req, _ := http.NewRequest(http.MethodPost, "/portal/chat/customer/000000000000000000000000/order/000000000000000000000000", nil)

			// serve request
			w := httptest.NewRecorder()
			tt.fields.router.Engine.ServeHTTP(w, req)

			if w.Code != tt.status {
				t.Fatalf("Expected to get status %v but instead got %v\n", tt.status, w.Code)
			}
		})
	}
}
