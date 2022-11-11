package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	configManager "gin-frame/lib/config"

	"github.com/rs/zerolog/log"
)

func main() {
	initSetting()

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("unknown error: %v", r)
			log.Fatal().Msgf("%s", err)
			time.Sleep(3 * time.Second)
		}
	}()
	handler, sqlDB, err := initializeService()
	defer func() {
		err := sqlDB.Close()
		if err != nil {
			log.Fatal().Err(err).Msgf("mysql db connection close error")
		}
	}()

	// start http server
	httpServer := &http.Server{
		Addr:    configManager.Global.Api.HTTPBind,
		Handler: handler.Gin,
	}
	go func() {
		// service connection
		log.Info().Msgf("listening and serving HTTP on %s", httpServer.Addr)
		err = httpServer.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Panic().Msgf("http server listen failed: %v", err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
	<-stopChan
	log.Info().Msgf("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Panic().Msgf("http server shutdown error: %v", err)
	} else {
		log.Info().Msgf("gracefully stopped")
	}
}
