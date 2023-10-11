package port

import (
	"context"
	"machship/internal/core/domain"
)

type ApiRequest interface {
	GetUserInfo(ctx context.Context, username string) (domain.UserData, error)
}
