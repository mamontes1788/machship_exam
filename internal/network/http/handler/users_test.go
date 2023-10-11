package handler

import (
	"context"
	"errors"
	"machship/internal/core/domain"
	"machship/internal/core/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetUsers(t *testing.T) {
	scenarios := []struct {
		testcase     string
		requestBody  string
		serviceError error
		httpStatus   int
	}{
		{
			testcase:     "Normal test",
			requestBody:  `{"usernames": ["user1", "user2", "user3"]}`,
			serviceError: nil,
			httpStatus:   http.StatusOK,
		},
		{
			testcase:     "Bad Request",
			requestBody:  ``,
			serviceError: nil,
			httpStatus:   http.StatusBadRequest,
		},
		{
			testcase:     "Service error",
			requestBody:  `{"usernames": ["user1", "user2", "user3"]}`,
			serviceError: errors.New("New Service Error"),
			httpStatus:   http.StatusInternalServerError,
		},
	}

	for _, tc := range scenarios {
		mockUserService := mock.NewMockUserService()
		mockUserService.GetUsersFunc = func(ctx context.Context, request domain.RetrieveUsersRequest) (domain.GetUserResponse, error) {
			return domain.GetUserResponse{}, tc.serviceError
		}

		userHandler := NewUsersHandler(mockUserService)

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/retrieveUsers", strings.NewReader(tc.requestBody))
		userHandler.GetUsers(c)

		if c.Writer.Status() != tc.httpStatus {
			t.Errorf("%s: expected status %d, got %d", tc.testcase, c.Writer.Status(), tc.httpStatus)
		}

	}

}
