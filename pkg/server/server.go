package server

import (
	"compress/flate"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"github.com/thepatrik/hellohttp/internal/handler"
)

// Option type
type Option func(*Config)

// Config type
type Config struct {
	Logger *log.Logger
	Port   int
}

// WithLogger sets a logger
func WithLogger(logger *log.Logger) Option {
	return func(cfg *Config) {
		cfg.Logger = logger
	}
}

// WithPort sets a port (defaults to 8080)
func WithPort(port int) Option {
	return func(cfg *Config) {
		cfg.Port = port
	}
}

// Server type
type Server struct {
	cfg        *Config
	httpServer *http.Server
}

// New creates a new server
func New(options ...Option) (*Server, error) {
	cfg := &Config{
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		Port:   8080,
	}
	for _, option := range options {
		option(cfg)
	}

	if cfg.Port < 1 {
		return nil, fmt.Errorf("invalid port number: %v", cfg.Port)
	}

	// Basic cors
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	})

	router := chi.NewRouter()
	router.Use(cors.Handler)
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Compress(flate.DefaultCompression))

	// Health check routes
	healthHandler := handler.Health()
	router.Handle("/health", healthHandler)

	strport := strconv.Itoa(cfg.Port)

	httpServer := &http.Server{
		Addr:    ":" + strport,
		Handler: router,
	}

	svr := &Server{
		cfg:        cfg,
		httpServer: httpServer,
	}

	return svr, nil
}

// Addr returns the address of the server
func (svr *Server) Addr() string {
	return svr.httpServer.Addr
}

// Start starts the server
func (svr *Server) Start() error {
	return svr.httpServer.ListenAndServe()
}

// Stop stops the server
func (svr *Server) Stop() error {
	svr.httpServer.SetKeepAlivesEnabled(false)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return svr.httpServer.Shutdown(ctx)
}
