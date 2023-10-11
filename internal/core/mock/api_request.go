package mock

import (
	"context"
	"machship/internal/core/domain"
)

type MockApiRequest struct {
	GetUserInfoFunc func(ctx context.Context, username string) (domain.UserData, error)
}

func (m *MockApiRequest) GetUserInfo(ctx context.Context, username string) (domain.UserData, error) {
	if m.GetUserInfoFunc != nil {
		return m.GetUserInfoFunc(ctx, username)
	}
	return domain.UserData{}, nil
}

func NewMockApiRequest() *MockApiRequest {
	return &MockApiRequest{}
}
