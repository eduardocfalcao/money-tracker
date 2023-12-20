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

	"github.com/eduardocfalcao/money-tracker/internal/healthcheck"
	"github.com/eduardocfalcao/money-tracker/internal/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	// authMiddleware "github.com/eduardocfalcao/money-tracker/server/auth/middleware"
)

type handlers struct {
	users users.Handlers
}

func main() {
	serverPort := 8080
	router := chi.NewRouter()

	// authenticationMiddleware := authMiddleware.CreateMiddleware()

	handlers := handlers{
		users: *users.NewHandler(),
	}

	router.Use(middleware.Logger)

	server := &http.Server{
		Addr:        fmt.Sprintf(":%d", serverPort),
		Handler:     router,
		ReadTimeout: 30 * time.Second,
	}

	registerRoutes(router, handlers)

	Start(server)
}

func registerRoutes(r *chi.Mux, handlers handlers) {
	r.Get("/healthcheck", healthcheck.Healthcheck)

	r.Get("/users/me", handlers.users.Me)
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
