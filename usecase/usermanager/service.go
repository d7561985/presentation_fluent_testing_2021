package usermanager

import (
	"presentation_fluent_testing_2021/repository"
	"presentation_fluent_testing_2021/usecase"
)

type service struct {
	repo repository.Sessions
}

func New(repo repository.Sessions) usecase.UserManager {
	return &service{repo: repo}
}
