package port

import (
	"context"
	"machship/internal/core/domain"
)

type UserService interface {
	GetUsers(ctx context.Context, request domain.RetrieveUsersRequest) (domain.GetUserResponse, error)
}
