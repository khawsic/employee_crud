package data

import (
	"context"
	"errors"

	"employee/internal/biz"

	"gorm.io/gorm"
)

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) biz.EmployeeRepo {
	return &employeeRepo{
		db: db,
	}
}

func (r *employeeRepo) Create(ctx context.Context, e *biz.Employee) (*biz.Employee, error) {

	model := &EmployeeModel{
		Name:  e.Name,
		Email: e.Email,
		Role:  e.Role,
	}

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("employee with this email already exists")
		}
		return nil, err
	}

	e.ID = model.ID
	e.CreatedAt = model.CreatedAt
	e.UpdatedAt = model.UpdatedAt

	return e, nil
}

func (r *employeeRepo) Get(ctx context.Context, id int64) (*biz.Employee, error) {

	var model EmployeeModel

	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("employee not found")
		}
		return nil, err
	}

	return &biz.Employee{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		Role:      model.Role,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

func (r *employeeRepo) List(ctx context.Context, offset, limit int) ([]*biz.Employee, error) {

	var models []EmployeeModel

	if err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Find(&models).Error; err != nil {
		return nil, err
	}

	employees := make([]*biz.Employee, 0, len(models))
	for _, m := range models {
		employees = append(employees, &biz.Employee{
			ID:        m.ID,
			Name:      m.Name,
			Email:     m.Email,
			Role:      m.Role,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		})
	}

	return employees, nil
}

func (r *employeeRepo) Update(ctx context.Context, e *biz.Employee) (*biz.Employee, error) {

	var model EmployeeModel

	if err := r.db.WithContext(ctx).First(&model, e.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("employee not found")
		}
		return nil, err
	}

	model.Name = e.Name
	model.Email = e.Email
	model.Role = e.Role

	if err := r.db.WithContext(ctx).Save(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("employee with this email already exists")
		}
		return nil, err
	}

	e.CreatedAt = model.CreatedAt
	e.UpdatedAt = model.UpdatedAt

	return e, nil
}

func (r *employeeRepo) Delete(ctx context.Context, id int64) error {

	result := r.db.WithContext(ctx).Delete(&EmployeeModel{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("employee not found")
	}

	return nil
}
