package main

// @title My Go API
// @version 1.0
// @description This is a sample server for a Go API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /
// @schemes http

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/Konstantin299/EduTodo.git/internal/rest"
	"github.com/Konstantin299/EduTodo.git/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	serv := service.New(log)
	server := rest.New(log, "localhost", "3000", serv)

	if err := server.Run(ctx); err != nil {
		log.Panic(err)
	}
}
