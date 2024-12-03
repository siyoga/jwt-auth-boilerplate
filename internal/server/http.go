package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/siyoga/jwt-auth-boilerplate/internal/config"
	"github.com/siyoga/jwt-auth-boilerplate/internal/handler"
	"github.com/siyoga/jwt-auth-boilerplate/internal/log"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type (
	Server interface {
		Start()
		Stop() error
	}

	server struct {
		log      log.Logger
		cfg      config.Base
		server   *http.Server
		wg       sync.WaitGroup
		listener net.Listener

		middleware  handler.Middleware
		authHandler handler.Handler
	}
)

func (s *server) Start() {
	s.log.Zap().Info("Start app server", zap.Int("port", s.cfg.ServerPort))

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		if err := s.server.Serve(s.listener); err != nil && err != http.ErrServerClosed {
			s.log.Zap().Panic("Error while server app server", zap.Error(err))
		}
	}()
}

func (s *server) Stop() error {
	s.log.Zap().Info("Stop app server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	s.wg.Wait()
	return nil
}

func NewHttpServer(
	log log.Logger,
	cfg config.Base,

	middleware handler.Middleware,
	authHandler handler.Handler,
) (Server, error) {
	var err error
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.ServerPort))
	if err != nil {
		return nil, fmt.Errorf("cannot listen app port: %w", err)
	}

	router := mux.NewRouter()
	server := &server{
		log: log.Named("app_server"),
		cfg: cfg,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
			Handler: router,
		},
		listener:    listener,
		middleware:  middleware,
		authHandler: authHandler,
	}
	server.initRoutes(router)
	return server, nil
}

func (s *server) initRoutes(router *mux.Router) {
	s.authHandler.FillHandlers(router)
}
