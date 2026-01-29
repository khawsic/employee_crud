package biz

import "context"

type EmployeeRepo interface {
	Create(ctx context.Context, e *Employee) (*Employee, error)
	Get(ctx context.Context, id int64) (*Employee, error)
	List(ctx context.Context, offset, limit int) ([]*Employee, error)

	Update(ctx context.Context, e *Employee) (*Employee, error)
	Delete(ctx context.Context, id int64) error
}
