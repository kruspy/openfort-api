package main

import (
	"fmt"
	"net/http"
	"openfort-api/api/router"
	"openfort-api/cmd/openfort-api/config"
	"openfort-api/cmd/openfort-api/logger"
	"time"
)

func init() {
	logger.Init()
	config.Load()
}

const (
	ReadTimeout  time.Duration = time.Second * 60
	WriteTimeout time.Duration = time.Second * 60
)

func main() {
	cfg := config.GetConfig()

	addr := fmt.Sprintf("%s:%s", cfg.Address, cfg.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      router.InitRouter(),
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
	}

	logger.Info(fmt.Sprintf("[HTTP] Listening on %s", addr))

	if err := server.ListenAndServe(); err != nil {
		logger.Error(err)
	}
}
