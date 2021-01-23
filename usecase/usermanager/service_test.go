// +build unit

package usermanager

import (
	"fmt"
	"presentation_fluent_testing_2021/repository/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	srv *service

	repo *mocks.Sessions
}

func (s *Suite) SetupSuite() {
	s.srv = New(nil).(*service)
}

func (s *Suite) TearDownSuite() {}

func (s *Suite) SetupTest() {
	fmt.Printf("NAME: %s\n", s.T().Name())

	// reset mock
	s.repo = new(mocks.Sessions)
	s.srv.repo = s.repo

	// perform some extra work per test
	switch s.T().Name() {
	case "TestUserManager/TestAuth":
		// do some specific for test
	}
}

func (s *Suite) TearDownTest() {}

func TestUserManager(t *testing.T) {
	suite.Run(t, new(Suite))
}
