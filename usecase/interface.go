package usecase

import (
	"context"
	"errors"
)

var (
	ErrNotExist = errors.New("not exist")
	ErrInternal = errors.New("internal error")
	ErrBadEmail = errors.New("bad email")
)

type UserManager interface {
	Auth(ctx context.Context, email string, data []byte) (uint, error)
}
