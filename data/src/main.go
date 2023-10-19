package main

import (
	"casbin-service/logger"
	"time"
)

func init() {
	logger.Init()
}
func main() {
	defer func() {
		logger.Close()
	}()

	logger.Warn("test warning")
	logger.Error("test error")
	logger.Debug("test debug")
	for {
		logger.Info("hello world")
		time.Sleep(1 * time.Second)
	}
}
