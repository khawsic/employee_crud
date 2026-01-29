package server

import (
	"employee/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(
	logger log.Logger,
	employeeService *service.EmployeeService,
) *khttp.Server {

	srv := khttp.NewServer()

	r := srv.Route("/")

	r.POST("/employees", employeeService.CreateEmployee)
	r.GET("/employees/{id}", employeeService.GetEmployee)
	r.GET("/employees", employeeService.ListEmployees)
	r.PUT("/employees/{id}", employeeService.UpdateEmployee)
	r.DELETE("/employees/{id}", employeeService.DeleteEmployee)

	return srv
}
