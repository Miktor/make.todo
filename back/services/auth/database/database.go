package database

import (
	"context"

	"github.com/Miktor/make.todo/back/cmd/auth/models"
)

type AuthDB interface {
	RegisterUser(ctx context.Context, user *models.UserInfo) error
	LoginUser(ctx context.Context, user *models.UserInfo) error

	Close()
}
