package router

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/LeandroFranciscato/go-chatbot/internal/delivery/rest"
	mocks "github.com/LeandroFranciscato/go-chatbot/mocks/usecase/customer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Test_router_home(t *testing.T) {
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
			status: http.StatusOK,
			mockFn: func(customerMock *mocks.Customer) {},
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

			// create temp html file
			fileName := "home.html"
			_, _ = os.Create(fileName)
			tt.fields.router.Engine.LoadHTMLFiles(fileName)
			defer func() {
				_ = os.Remove(fileName)
			}()

			// add routes
			tt.fields.router.home()

			// prepare request
			req, _ := http.NewRequest(http.MethodGet, "/home", nil)

			// serve request
			w := httptest.NewRecorder()
			tt.fields.router.Engine.ServeHTTP(w, req)

			if w.Code != tt.status {
				t.Fatalf("Expected to get status %v but instead got %v\n", tt.status, w.Code)
			}
		})
	}
}
