package main

import (
	"employee_management/config"
	"employee_management/internal/server"
	"employee_management/internal/server/handler"
	"employee_management/pkg/employee"
	"employee_management/utils/initialize"

	"go.uber.org/fx"
)

func serverRun() {
	app := fx.New(
		fx.Provide(
			// postgresql
			initialize.NewDB,
		),
		config.Module,
		initialize.Module,
		server.Module,
		handler.Module,
		employee.Module,
	)

	// Run app forever
	app.Run()
}
