package router

import (
	"net/http"
	"net/http/httptest"
	"review-chatbot/internal/delivery/rest"
	"review-chatbot/internal/domain/entity"
	mocks "review-chatbot/mocks/usecase/customer"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func Test_router_login(t *testing.T) {
	type fields struct {
		router router
	}
	tests := []struct {
		name   string
		fields fields
		status int
		mockFn func(mock *mocks.Customer)
	}{
		{
			name: "Success",
			fields: fields{
				router{
					Server: rest.Server{
						Customer: &mocks.Customer{},
					},
				},
			},
			status: http.StatusPermanentRedirect,
			mockFn: func(customerMock *mocks.Customer) {
				customerMock.EXPECT().Login(mock.Anything, "email", "1a1dc91c907325c69271ddf0c944bc72").Return(entity.Customer{Email: "email"}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// run mock
			tt.mockFn(tt.fields.router.Server.Customer.(*mocks.Customer))

			// prepare router
			tt.fields.router.Engine = gin.Default()
			store := cookie.NewStore([]byte("my-secret-key"))
			tt.fields.router.Engine.Use(sessions.Sessions("session", store))

			// add routes
			tt.fields.router.login()

			// prepare request
			req, err := http.NewRequest(http.MethodPost, "/login", nil)
			if err != nil {
				t.Fatalf("Couldn't create request: %v\n", err)
			}

			// set parameters
			req.Form = make(map[string][]string)
			req.Form.Add("email", "email")
			req.Form.Add("password", "pass")

			// serve request
			w := httptest.NewRecorder()
			tt.fields.router.Engine.ServeHTTP(w, req)

			if w.Code != tt.status {
				t.Fatalf("Expected to get status %v but instead got %v\n", http.StatusOK, w.Code)
			}
		})
	}
}
