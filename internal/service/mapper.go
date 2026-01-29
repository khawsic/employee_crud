package service

import (
	"employee/internal/biz"
	"employee/internal/service/dto"
)

func ToEmployeeResponse(e *biz.Employee) *dto.EmployeeResponse {
	if e == nil {
		return nil
	}
	return &dto.EmployeeResponse{
		ID:        e.ID,
		Name:      e.Name,
		Email:     e.Email,
		Role:      e.Role,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func ToEmployeeResponseList(employees []*biz.Employee) []*dto.EmployeeResponse {
	resp := make([]*dto.EmployeeResponse, 0, len(employees))
	for _, e := range employees {
		resp = append(resp, ToEmployeeResponse(e))
	}
	return resp
}

func FromCreateRequest(req *dto.CreateEmployeeRequest) *biz.Employee {
	return &biz.Employee{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}
}

func FromUpdateRequest(req *dto.UpdateEmployeeRequest) *biz.Employee {
	return &biz.Employee{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}
}
