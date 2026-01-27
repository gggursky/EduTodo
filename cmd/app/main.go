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
	"os"
	"os/signal"
	"syscall"

	"github.com/Konstantin299/EduTodo.git/internal/rest"
	"github.com/Konstantin299/EduTodo.git/internal/service"
	"github.com/Konstantin299/EduTodo.git/internal/store"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

const dsnValue = "postgres://user:secret@localhost:5533/edu_todo_db?sslmode=disable"

func main() {

	dsn := getEnv(
		"AS_PG_DSN",
		dsnValue)

	log := logrus.New()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	db, err := store.New(ctx, log, dsn)
	if err != nil {
		log.Panicf("store.New(ctx, log, dsn): %v", err)
	}

	if err = db.Migrate(migrate.Up); err != nil {
		log.Panicf("db.Migrate(migrate.Up): %v", err)
	}

	serv := service.New(log, db)
	server := rest.New(log, "localhost", "3000", serv)

	if err := server.Run(ctx); err != nil {
		log.Panic(err)
	}
}

func getEnv(env, defaultValue string) string {
	val := os.Getenv(env)
	if val == "" {
		val = defaultValue
	}

	return val
}
