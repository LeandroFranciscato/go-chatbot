package router

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/LeandroFranciscato/go-chatbot/internal/delivery/rest"
	"github.com/LeandroFranciscato/go-chatbot/internal/domain/entity"
	chatmocks "github.com/LeandroFranciscato/go-chatbot/mocks/usecase/chat"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_router_chat_list(t *testing.T) {
	type fields struct {
		router router
	}
	tests := []struct {
		name   string
		fields fields
		status int
		mockFn func(chatmock *chatmocks.Chat)
	}{
		{
			name: "Success",
			fields: fields{
				router{
					Server: rest.Server{
						Chat: &chatmocks.Chat{},
					},
				},
			},
			status: http.StatusOK,
			mockFn: func(chatmock *chatmocks.Chat) {
				objID, _ := primitive.ObjectIDFromHex("000000000000000000000000")
				chatmock.EXPECT().FindByCustomer(mock.Anything, objID).Return([]entity.Chat{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// run mock
			tt.mockFn(tt.fields.router.Server.Chat.(*chatmocks.Chat))

			// prepare router
			tt.fields.router.Engine = gin.Default()

			// create temp html file
			fileName := "chat_list.html"
			_, _ = os.Create(fileName)
			tt.fields.router.Engine.LoadHTMLFiles(fileName)
			defer func() {
				_ = os.Remove(fileName)
			}()

			// add routes
			tt.fields.router.Engine.Handle(http.MethodGet, "/portal/chat/list", func(ctx *gin.Context) {
				tt.fields.router.chatListHandler(ctx, "000000000000000000000000")
			})

			// prepare request
			req, _ := http.NewRequest(http.MethodGet, "/portal/chat/list", nil)

			// serve request
			w := httptest.NewRecorder()
			tt.fields.router.Engine.ServeHTTP(w, req)

			if w.Code != tt.status {
				t.Fatalf("Expected to get status %v but instead got %v\n", tt.status, w.Code)
			}
		})
	}
}
