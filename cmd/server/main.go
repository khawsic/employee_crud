package main

import (
	"os"

	"employee/internal/biz"
	"employee/internal/data"
	"employee/internal/server"
	"employee/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

func main() {
	logger := log.NewStdLogger(os.Stdout)

	db, err := data.NewDB(
		"host=127.0.0.1 user=postgres password=root dbname=Employee port=5432 sslmode=disable TimeZone=Asia/Kolkata",
	)
	if err != nil {
		panic(err)
	}

	employeeRepo := data.NewEmployeeRepo(db)
	employeeUsecase := biz.NewEmployeeUseCase(employeeRepo)
	employeeService := service.NewEmployeeService(employeeUsecase)

	httpServer := server.NewHTTPServer(logger, employeeService)

	app := kratos.New(
		kratos.Name("employee-service"),
		kratos.Logger(logger),
		kratos.Server(httpServer),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
