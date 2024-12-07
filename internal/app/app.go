package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	logger "github.com/mizmorr/loggerm"
	"github.com/mizmorr/wallet/config"
	"github.com/mizmorr/wallet/internal/controller"
	"github.com/mizmorr/wallet/internal/router"
	"github.com/mizmorr/wallet/internal/service"
	"github.com/mizmorr/wallet/pkg/server"
	"github.com/mizmorr/wallet/pkg/types"
	"github.com/mizmorr/wallet/store"
)

const loggerKey types.ContextKey = "logger"

func Run() error {
	// Inintialize
	conf := config.Get()

	log := logger.Get(conf.LogLevel)

	ctx := context.WithValue(context.Background(), loggerKey, log)

	interrupt := make(chan os.Signal, 1)

	log.Debug().Msg("[app.Run] - store initialization...")
	store, err := store.New(ctx)
	if err != nil {
		return err
	}

	log.Debug().Msg("[app.Run] - service for songs initialization...")
	ws, err := service.NewWalletService(store, ctx)
	if err != nil {
		return err
	}

	log.Debug().Msg("[app.Run] - controller initialization...")
	controller := controller.NewWalletController(ws, ctx)

	log.Debug().Msg("[app.Run] - starting server...")
	handler := gin.New()
	router.NewRouter(handler, controller)

	httpServer := server.New(handler, conf.HTTPAddr)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	log.Info().Msg("Server is running ...")

	select {
	case s := <-interrupt:
		log.Info().Msg("[app.Run] - signal " + s.String())
		time.Sleep(500 * time.Millisecond)
	case err = <-httpServer.Notify():
		log.Error().Err(fmt.Errorf("[app.Run] - httpServer.Notify " + err.Error()))
	}

	return httpServer.Shutdown()
}
