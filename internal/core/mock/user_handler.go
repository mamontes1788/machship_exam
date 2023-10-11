package mock

import (
	"context"
	"machship/internal/core/domain"
)

type MockUserService struct {
	GetUsersFunc func(ctx context.Context, request domain.RetrieveUsersRequest) (domain.GetUserResponse, error)
}

func (m *MockUserService) GetUsers(ctx context.Context, request domain.RetrieveUsersRequest) (domain.GetUserResponse, error) {
	if m.GetUsersFunc != nil {
		return m.GetUsersFunc(ctx, request)
	}
	return domain.GetUserResponse{}, nil
}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}
