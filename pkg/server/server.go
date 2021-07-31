package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atamano/fizz-buzz/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Config for server
type Config struct {
	Port int
}

//Server API
type Server struct {
	Router *gin.Engine
	config Config
}

type HandlerFunc func(interface{})

type Router = gin.IRouter
type Context = gin.Context

func secureHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Protects from MimeType Sniffing
		c.Header("X-Content-Type-Options", "nosniff")
		// Prevents browser from prefetching DNS
		c.Header("X-DNS-Prefetch-Control", "off")
		// Denies website content to be served in an iframe
		c.Header("X-Frame-Options", "DENY")
		c.Header("Strict-Transport-Security", "max-age=5184000; includeSubDomains")
		// Prevents Internet Explorer from executing downloads in site's context
		c.Header("X-Download-Options", "noopen")
		// Minimal XSS protection
		c.Header("X-XSS-Protection", "1; mode=block")
	}
}

//NewGroup returns a new group
func (s *Server) NewGroup(groupPrefix string) Router {
	return s.Router.Group(groupPrefix)
}

//Run server
func (s *Server) Run() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.Port),
		Handler: s.Router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listenAndServe failed", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown failed", err.Error())
	}

	logger.Info("Server exiting")
}

//New server
func New(config Config) Server {
	router := gin.Default()

	router.Use(cors.Default())
	router.Use(secureHeaders())

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"message": "ok",
		})
	})

	return Server{router, config}
}
