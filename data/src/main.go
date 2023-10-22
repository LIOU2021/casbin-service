package main

import (
	"casbin-service/config"
	"casbin-service/logger"
	"casbin-service/router"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	config.Init()
	logger.Init()
}

func main() {
	defer func() {
		logger.Close()
	}()

	r := router.Init()
	p := fmt.Sprintf(":%s", config.Config.Server.Port)
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
