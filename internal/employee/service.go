package employee

import (
	"context"
)

type Service interface {
	GetActiveSmithEmployees(ctx context.Context) ([]EmployeeName, error)
	GetEmployeesWithoutReviews(ctx context.Context) ([]EmployeeName, error)
	GetHireDateDiffActiveEmployees(ctx context.Context) (int, error)
	GetSalaryEstimationWithReviews(ctx context.Context) ([]EmployeeSalaryEstimate, error)
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

func (s *service) GetEmployeesWithoutReviews(ctx context.Context) ([]EmployeeName, error) {
	return s.repo.FindEmployeesWithoutReviews(ctx)
}

func (s *service) GetHireDateDiffActiveEmployees(ctx context.Context) (int, error) {
	return s.repo.GetHireDateDiffActiveEmployees(ctx)
}

func (s *service) GetSalaryEstimationWithReviews(ctx context.Context) ([]EmployeeSalaryEstimate, error) {
	return s.repo.GetSalaryEstimationWithReviews(ctx)
}
