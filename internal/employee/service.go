package employee

import (
	"context"
)

type Service interface {
	GetActiveSmithEmployees(ctx context.Context) ([]EmployeeName, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetActiveSmithEmployees(ctx context.Context) ([]EmployeeName, error) {
	return s.repo.FindActiveSmithEmployees(ctx)
}
