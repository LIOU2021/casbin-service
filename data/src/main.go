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
)

func init() {
	logger.Init()
}
func main() {
	defer func() {
		logger.Close()
	}()

	r := router.Init()

	p := ":8080"
	srv := &http.Server{
		Addr:    p,
		Handler: r,
	}
	go func() {
		logger.Infof("start listen %s", p)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("ListenAndServe | err: %v", err)
			os.Exit(1)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	logger.Infof("shutdown")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(c); err != nil {
		logger.Errorf("Shutdown | err: %v", err)
		os.Exit(1)
		return
	}
}
