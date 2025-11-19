package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Konstantin299/EduTodo.git/internal/models"
	_ "github.com/Konstantin299/EduTodo.git/internal/swagger/docs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type service interface {
	GetInfo() (string, error)
	GetAnswersList() ([]models.ResponseQuestion, error)
	CheckAnswer([]models.UserAnswer) ([]models.CheckResult, error)
}

type Server struct {
	log             *logrus.Entry
	host            string
	port            string
	server          *http.Server
	shutdownTimeout time.Duration
	service         service
}

func New(log *logrus.Logger, host string, port string, service service) *Server {
	return &Server{
		log:     log.WithField("module", "rest"),
		host:    host,
		port:    port,
		service: service,
	}
}

func (s *Server) Run(ctx context.Context) error {
	router := gin.Default()

	runServerFn := func() error {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("running http server: %w", err)
		}

		return nil
	}

	go func() {
		<-ctx.Done()

		gfCtx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
		defer cancel()

		if err := s.server.Shutdown(gfCtx); err != nil { //nolint:contextcheck
			s.log.Warnf("s.sever.Shutdown(gfCtx): %v", err)
		}
	}()
	// маршрут.
	router.GET("/text", s.GetText)
	router.GET("/answers", s.GetAnswers)
	router.GET("/check/:answerCode/:respCode", s.Check)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/check", s.Check)

	s.server = &http.Server{
		Addr:              fmt.Sprintf("%s:%s", s.host, s.port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second, //nolint:gomnd
	}

	s.log.Infof("starting  server on %s", s.server.Addr)

	return runServerFn()
}
