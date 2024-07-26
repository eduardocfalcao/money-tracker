package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eduardocfalcao/money-tracker/internal/container"
	"github.com/eduardocfalcao/money-tracker/internal/healthcheck"
	"github.com/eduardocfalcao/money-tracker/internal/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type router interface {
	RegisterRoutes(router chi.Router, privateRoutes chi.Router)
}

func main() {
	serverPort := 8080
	router := chi.NewRouter()
	secretKey := "secret-key" // load the key from somewhere
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:12345678a@localhost:5433/money-tracker?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	c, err := container.NewContainer(secretKey, conn)
	if err != nil {
		logrus.Errorf("Error creating the container to start the application: %s", err)
		return
	}

	router.Use(middleware.Logger)

	server := &http.Server{
		Addr:        fmt.Sprintf(":%d", serverPort),
		Handler:     router,
		ReadTimeout: 30 * time.Second,
	}

	privateRouter := router.With(c.JWTMiddlewareService.VerifyTokenMiddleware)
	registerRoutes(router, privateRouter, c)

	Start(server)
}

func registerRoutes(r chi.Router, privateRouter chi.Router, c *container.Container) {
	// public routes
	r.Get("/healthcheck", healthcheck.Healthcheck)
	routers := []router{
		users.NewRoutes(c.UsersHanders),
	}
	// private routes
	for _, router := range routers {
		router.RegisterRoutes(r, privateRouter)
	}
}

func Start(server *http.Server) {
	go func() {
		logrus.Infof("# Starting server %s", server.Addr)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stopCh

	logrus.Info("Terminating service...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.Error(err)
	}
	logrus.Info("# Service terminated")
}
