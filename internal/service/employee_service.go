package service

import (
	"net/http"
	"strconv"

	"employee/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

type EmployeeService struct {
	uc *biz.EmployeeUsecase
}

func NewEmployeeService(uc *biz.EmployeeUsecase) *EmployeeService {
	return &EmployeeService{uc: uc}
}

func (s *EmployeeService) CreateEmployee(ctx khttp.Context) error {
	var emp biz.Employee

	if err := ctx.Bind(&emp); err != nil {
		return errors.BadRequest("INVALID_BODY", err.Error())
	}

	res, err := s.uc.CreateEmployee(ctx, &emp)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *EmployeeService) GetEmployee(ctx khttp.Context) error {
	vars := ctx.Vars()

	idStr, ok := vars["id"]
	if !ok {
		return errors.BadRequest("MISSING_ID", "employee id is required")
	}

	id, err := strconv.ParseInt(idStr[0], 10, 64)
	if err != nil {
		return errors.BadRequest("INVALID_ID", "employee id must be a number")
	}

	res, err := s.uc.GetEmployee(ctx, id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (s *EmployeeService) ListEmployees(ctx khttp.Context) error {

	offset, _ := strconv.Atoi(ctx.Query().Get("offset"))
	limit, _ := strconv.Atoi(ctx.Query().Get("limit"))

	res, err := s.uc.ListEmployees(ctx, offset, limit)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (s *EmployeeService) UpdateEmployee(ctx khttp.Context) error {
	var emp biz.Employee

	vars := ctx.Vars()
	idStr, ok := vars["id"]
	if !ok {
		return errors.BadRequest("MISSING_ID", "employee id is required")
	}

	id, err := strconv.ParseInt(idStr[0], 10, 64)
	if err != nil {
		return errors.BadRequest("INVALID_ID", "employee id must be a number")
	}

	if err := ctx.Bind(&emp); err != nil {
		return errors.BadRequest("INVALID_BODY", err.Error())
	}

	emp.ID = id

	res, err := s.uc.UpdateEmployee(ctx, &emp)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (s *EmployeeService) DeleteEmployee(ctx khttp.Context) error {

	vars := ctx.Vars()
	idStr, ok := vars["id"]
	if !ok {
		return errors.BadRequest("MISSING_ID", "employee id is required")
	}

	id, err := strconv.ParseInt(idStr[0], 10, 64)
	if err != nil {
		return errors.BadRequest("INVALID_ID", "employee id must be a number")
	}

	if err := s.uc.DeleteEmployee(ctx, id); err != nil {
		return err
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
