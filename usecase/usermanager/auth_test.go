// +build unit

package usermanager

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"presentation_fluent_testing_2021/models"
	"presentation_fluent_testing_2021/usecase"

	"github.com/jackc/fake"
	"github.com/stretchr/testify/mock"
)

// TestAuth ID: TestUserManager/TestAuth
func (s *Suite) TestAuth() {
	type req struct {
		email string
		data  []byte
	}

	type res struct {
		id  uint
		err error
	}

	tests := []struct {
		name string
		req  req
		res  res

		prepare func(req, res)
	}{
		{
			"OK",

			req{fake.EmailAddress(), []byte{'H', 'E', 'L', 'O'}},
			res{uint(math.MaxUint32), nil},

			func(r req, res res) {
				usr := models.User{
					ID:    uint(rand.Uint64()),
					Name:  fake.FullName(),
					Email: r.email,
				}
				ses := models.Session{
					ID:     res.id,
					UserID: usr.ID,
					BLOB:   []byte(fake.Brand()),
				}

				s.repo.On("Authorize", mock.Anything, r.email, r.data).Return(usr, nil)
				s.repo.On("CreateSession", mock.Anything, usr).Return(ses, nil)
			},
		},
		{
			"empty email",
			req{"", []byte("H")},
			res{0, usecase.ErrBadEmail},
			func(req req, res res) {},
		},
		{
			"bad auth",
			req{fake.EmailAddress(), []byte{'H', 'E', 'L', 'O'}},
			res{0, usecase.ErrNotExist},
			func(req req, res res) {
				s.repo.On("Authorize", mock.Anything, req.email, req.data).
					Return(models.User{}, res.err)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.prepare(test.req, test.res)

			id, err := s.srv.Auth(context.Background(), test.req.email, test.req.data)
			fmt.Println(errors.Is(err, test.res.err))
			s.Equal(test.res.id, id)
		})
	}
}
