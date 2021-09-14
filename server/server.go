package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sharpvik/log-go/v2"

	"github.com/lisn-rocks/lisn/configs"
	"github.com/lisn-rocks/lisn/database"
)

type Server struct {
	*echo.Echo
}

func New() (s *Server) {
	_ = database.Init()
	return &Server{runtime()}
}

func (s *Server) Grace(done chan bool) {
	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Debug("stopping server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("graceful server shutdown failed: %s", err)
	} else {
		log.Debug("server stopped")
	}

	close(done)
}

func (s *Server) Serve() {
	log.Debugf("serving at %s ...", configs.Server.Address)
	if err := s.Start(configs.Server.Address); err != http.ErrServerClosed {
		log.Errorf("server shut with error: %s", err)
	}
}

func (s *Server) ServeWithGrace(done chan bool) {
	go s.Grace(done)
	go s.Serve()
}
