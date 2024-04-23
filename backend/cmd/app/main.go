package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/config"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/server"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	transport "github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/auth"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/database/postgres"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/logger"
	"github.com/subosito/gotenv"
)

func main() {
	//* Init config
	if err := gotenv.Load("../.env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	conf, err := config.Init("configs/config.yaml")
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	logger.NewLogger(logger.WithLevel(conf.LogLevel), logger.WithAddSource(conf.LogSource))

	//* Dependencies
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     conf.Postgres.Host,
		Port:     conf.Postgres.Port,
		Username: conf.Postgres.Username,
		Password: conf.Postgres.Password,
		DBName:   conf.Postgres.DbName,
		SSLMode:  conf.Postgres.SSLMode,
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	keycloak := auth.NewKeycloakClient(auth.Deps{
		Url:       conf.Keycloak.Url,
		ClientId:  conf.Keycloak.ClientId,
		Realm:     conf.Keycloak.Realm,
		AdminName: conf.Keycloak.Root,
		AdminPass: conf.Keycloak.RootPass,
	})

	//* Services, Repos & API Handlers
	repos := repository.NewRepository(db)
	services := services.NewServices(services.Deps{
		Repos:           repos,
		Keycloak:        keycloak,
		AccessTokenTTL:  conf.Auth.AccessTokenTTL,
		RefreshTokenTTL: conf.Auth.RefreshTokenTTL,
		ErrorBotUrl:     conf.ErrorBot.Url,
		BotUrl:          conf.Bot.Url,
	})

	handlers := transport.NewHandler(services, keycloak)

	//* HTTP Server

	// if err := services.Notification.Start(&conf.Notification); err != nil {
	// 	logger.Fatalf("failed to start sending notification. error: %s\n", err.Error())
	// }

	srv := server.NewServer(conf, handlers.Init(conf))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	logger.Info("Application started on port: %s", conf.Http.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	// if err := services.Notification.Stop(); err != nil {
	// 	logger.Errorf("failed to stop sending notification. error: %s", err)
	// }

	if err := srv.Stop(ctx); err != nil {
		logger.Error("failed to stop server:", logger.ErrAttr(err))
	}
}
