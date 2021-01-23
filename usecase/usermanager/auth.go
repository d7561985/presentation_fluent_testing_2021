package usermanager

import (
	"context"
	"presentation_fluent_testing_2021/usecase"
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (s *service) Auth(ctx context.Context, email string, data []byte) (uint, error) {
	if !emailRegex.Match([]byte(email)) {
		return 0, usecase.ErrBadEmail
	}

	user, err := s.repo.Authorize(ctx, email, data)
	if err != nil {
		return 0, usecase.ErrNotExist
	}

	ses, err := s.repo.CreateSession(ctx, user)
	if err != nil {
		return 0, usecase.ErrInternal
	}

	return ses.ID, nil
}
