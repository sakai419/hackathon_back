package services

import "local-test/internal/repositories"

type Service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		repo: repo,
	}
}