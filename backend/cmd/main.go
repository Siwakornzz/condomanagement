package main

import (
	"condomanagement/internal/config"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {

	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// http framework
	app := fiber.New()

	// start server
	go func() {
		if err := app.Listen(cfg.App.GetAddress()); err != nil && err != http.ErrServerClosed {
			log.Error("[ERROR] shutting down the server", err)
		}
	}()

	log.Infof("%s started at %s", cfg.App.Name, cfg.App.GetAddress())

	// setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), cfg.App.GetTimeout())
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatal("shutting down the server:", map[string]interface{}{"error": err})

	}

}
