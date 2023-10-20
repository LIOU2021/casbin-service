package main

import (
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
	logger.Init()
}
func main() {
	defer func() {
		logger.Close()
	}()

	r := router.New()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	logger.Info("casbin-service shutdown ...")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(c); err != nil {
		logger.Errorf("casbin-service Shutdown error: %v", err)
		os.Exit(1)
		return
	}
}
