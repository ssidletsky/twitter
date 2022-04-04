package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/ssidletsky/esportal-twitter/app/config"
	"github.com/ssidletsky/esportal-twitter/app/tweets/delivery/http"
	"github.com/ssidletsky/esportal-twitter/app/tweets/repository/mysql"
	"github.com/ssidletsky/esportal-twitter/app/tweets/usecases"
)

// Configure configures application
func Configure(cnf *config.App) {
	config.SetLoggerConfig(cnf.Logger)
	mysql.Initialize(cnf.MySQL)
}

func main() {
	cnf, err := config.Get()
	if err != nil {
		log.Fatalf("config.Get error: %v", err)
	}
	Configure(cnf)

	app := fiber.New()
	tuc := usecases.NewTweet(mysql.NewRepository())
	http.RegisterTweetsController(app, tuc)

	go func() {
		log.Infof("Starting API on %s port", cnf.Server.Port)
		if err := app.Listen(fmt.Sprintf(":%s", cnf.Server.Port)); err != nil {
			log.Fatal("Server error: %w", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info("Shutting down server")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}
}
