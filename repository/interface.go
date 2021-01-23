package repository

import (
	"context"
	"presentation_fluent_testing_2021/models"
)

//go:generate mockery --all

type Sessions interface {
	Authorize(ctx context.Context, email string, token []byte) (models.User, error)

	CreateSession(context.Context, models.User) (models.Session, error)
	FindSession(context.Context, models.User) models.Session
}
