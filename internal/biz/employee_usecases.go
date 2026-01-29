package biz

import (
	"context"
	"errors"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type EmployeeUsecase struct {
	repo EmployeeRepo
}

func NewEmployeeUseCase(repo EmployeeRepo) *EmployeeUsecase {
	return &EmployeeUsecase{
		repo: repo,
	}
}

func (uc *EmployeeUsecase) CreateEmployee(ctx context.Context, e *Employee) (*Employee, error) {

	if e.Name == "" {
		return nil, errors.New("employee name is required")
	}
	if len(e.Name) < 2 || len(e.Name) > 100 {
		return nil, errors.New("employee name must be between 2 and 100 characters")
	}

	if e.Email == "" {
		return nil, errors.New("employee email is required")
	}
	if !emailRegex.MatchString(e.Email) {
		return nil, errors.New("invalid email format")
	}

	if e.Role == "" {
		return nil, errors.New("employee role is required")
	}

	return uc.repo.Create(ctx, e)
}

func (uc *EmployeeUsecase) GetEmployee(ctx context.Context, id int64) (*Employee, error) {

	if id <= 0 {
		return nil, errors.New("invalid employee id")
	}

	return uc.repo.Get(ctx, id)
}

func (uc *EmployeeUsecase) ListEmployees(ctx context.Context, offset, limit int) ([]*Employee, error) {

	if offset < 0 {
		return nil, errors.New("offset cannot be negative")
	}
	if limit <= 0 {
		return nil, errors.New("limit must be greater than zero")
	}

	return uc.repo.List(ctx, offset, limit)
}

func (uc *EmployeeUsecase) UpdateEmployee(ctx context.Context, e *Employee) (*Employee, error) {

	if e.ID <= 0 {
		return nil, errors.New("invalid employee id")
	}

	if e.Name == "" {
		return nil, errors.New("employee name is required")
	}
	if len(e.Name) < 2 || len(e.Name) > 100 {
		return nil, errors.New("employee name must be between 2 and 100 characters")
	}

	if e.Email == "" {
		return nil, errors.New("employee email is required")
	}
	if !emailRegex.MatchString(e.Email) {
		return nil, errors.New("invalid email format")
	}

	if e.Role == "" {
		return nil, errors.New("employee role is required")
	}

	return uc.repo.Update(ctx, e)
}

func (uc *EmployeeUsecase) DeleteEmployee(ctx context.Context, id int64) error {

	if id <= 0 {
		return errors.New("invalid employee id")
	}

	return uc.repo.Delete(ctx, id)
}
