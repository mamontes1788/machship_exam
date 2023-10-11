package http

import (
	"context"
	"machship/internal/core/util"
	"machship/internal/network/http/route"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(genericRoute route.Route) *ServerHTTP {
	engine := gin.Default()

	api := engine.Group("/api")
	{
		genericRoute.PublicSetup(api)
	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           sh.engine,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			util.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	util.Infoln("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		util.Fatalf("Server forced to shutdown: %v", err)
	}

	util.Infoln("Server exiting")
}
