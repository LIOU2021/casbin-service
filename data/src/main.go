package main

import (
	"casbin-service/logger"
	"casbin-service/router"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func init() {
	logger.Init()
}
func main() {
	defer func() {
		logger.Close()
	}()

	r := router.New()

	p := ":8080"
	srv := &http.Server{
		Addr:    p,
		Handler: r,
	}
	go func() {
		logger.Infof("start listen %s", p)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("ListenAndServe", zap.String("err", err.Error()))
			os.Exit(1)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	logger.Info("shutdown")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(c); err != nil {
		logger.Error("Shutdown", zap.String("err", err.Error()))
		os.Exit(1)
		return
	}
}
