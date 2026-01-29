package service

import (
	"net/http"
	"strconv"

	"employee/internal/biz"
	"employee/internal/service/dto"

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
	var req dto.CreateEmployeeRequest
	if err := ctx.Bind(&req); err != nil {
		return errors.BadRequest("INVALID_BODY", err.Error())
	}

	employee := FromCreateRequest(&req)
	created, err := s.uc.CreateEmployee(ctx, employee)
	if err != nil {
		switch err.Error() {
		case "employee name is required", "employee email is required", "employee role is required", "invalid email format":
			return errors.BadRequest("INVALID_INPUT", err.Error())
		case "employee with this email already exists":
			return errors.Conflict("DUPLICATE_EMPLOYEE", err.Error())
		default:
			return errors.New(500, "SERVER_ERROR", err.Error())

		}
	}

	resp := ToEmployeeResponse(created)
	return ctx.JSON(http.StatusCreated, resp)
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

	employee, err := s.uc.GetEmployee(ctx, id)
	if err != nil {
		switch err.Error() {
		case "employee not found":
			return errors.NotFound("EMPLOYEE_NOT_FOUND", err.Error())
		default:
			return errors.New(500, "SERVER_ERROR", err.Error())

		}
	}

	resp := ToEmployeeResponse(employee)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *EmployeeService) ListEmployees(ctx khttp.Context) error {
	// Parse query params
	page, _ := strconv.Atoi(ctx.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	size, _ := strconv.Atoi(ctx.Query().Get("size"))
	if size < 1 {
		size = 10
	}

	offset := (page - 1) * size

	employees, err := s.uc.ListEmployees(ctx, offset, size)
	if err != nil {
		return errors.New(500, "SERVER_ERROR", err.Error())

	}

	resp := ToEmployeeResponseList(employees)

	// Include pagination metadata
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": resp,
		"meta": map[string]interface{}{
			"page":  page,
			"size":  size,
			"total": len(employees), // in real projects, fetch total count from DB
		},
	})
}

func (s *EmployeeService) UpdateEmployee(ctx khttp.Context) error {
	var req dto.UpdateEmployeeRequest
	if err := ctx.Bind(&req); err != nil {
		return errors.BadRequest("INVALID_BODY", err.Error())
	}

	vars := ctx.Vars()
	idStr, ok := vars["id"]
	if !ok {
		return errors.BadRequest("MISSING_ID", "employee id is required")
	}

	id, err := strconv.ParseInt(idStr[0], 10, 64)
	if err != nil {
		return errors.BadRequest("INVALID_ID", "employee id must be a number")
	}

	employee := FromUpdateRequest(&req)
	employee.ID = id

	updated, err := s.uc.UpdateEmployee(ctx, employee)
	if err != nil {
		switch err.Error() {
		case "employee not found":
			return errors.NotFound("EMPLOYEE_NOT_FOUND", err.Error())
		case "employee with this email already exists":
			return errors.Conflict("DUPLICATE_EMPLOYEE", err.Error())
		default:
			return errors.New(500, "SERVER_ERROR", err.Error())

		}
	}

	resp := ToEmployeeResponse(updated)
	return ctx.JSON(http.StatusOK, resp)
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
		switch err.Error() {
		case "employee not found":
			return errors.NotFound("EMPLOYEE_NOT_FOUND", err.Error())
		default:
			return errors.New(500, "SERVER_ERROR", err.Error())

		}
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
