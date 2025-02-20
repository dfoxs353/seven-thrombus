package main

import (
	"context"
	"fmt"
	"log/slog"
	"main/http/swagger"
	"main/internal/files"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/cors"
	http_swagger "github.com/swaggo/http-swagger"
	config "gitlab.com/volgaIt/packages/config"
	mw "gitlab.com/volgaIt/packages/middleware"
	"gitlab.com/volgaIt/packages/postgres"
	server "gitlab.com/volgaIt/packages/server"
)

type cfg struct {
	Postgres config.Postgres `mapstructure:"postgres"`
	Server   config.Server   `mapstructure:"server"`
	Minio    files.MinioCfg  `mapstructure:"minio"`
}

const ServiceName = "users"

// @title           Accounts MS
// @version         1.0
// @BasePath  /
// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
//
//go:generate swag init --parseDependency --parseInternal -g ./main.go -o ./http/swagger
func main() {
	// инициализация конфига
	cfg, err := config.ReadConfig[cfg](os.Getenv("CONFIG_PATH"))
	if err != nil {
		slog.Error("failed to read config", "err", err)
	}

	cfg.Server.ShutdownTimeout *= time.Second
	cfg.Tokens.AccessTTL *= time.Second
	cfg.Tokens.RefreshTTL *= time.Second

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, os.Interrupt, os.Kill)
	defer cancel()

	// инициализация требуемых зависимостей
	db, err := postgres.ConnectWithPing(
		ctx,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)
	if err != nil {
		slog.Error("failed connect to postgres", "err", err)
		return
	}
	defer db.Close()

	var (
	// files handlers
	)

	// инициализация http сервера
	srv := server.New(cfg.Server.Host, cfg.Server.Port, cfg.Server.ShutdownTimeout)

	srv.Post("/api/upload", mw.WrapErrDetail(nil))
	srv.Get("/api/meta/{id}", mw.WrapErrDetail(nil))
	srv.Get("/api/download/{id}", mw.WrapErrDetail(nil))

	swagger.SwaggerInfo.BasePath = fmt.Sprintf("/%s/", ServiceName)
	srv.Get("/swagger/{*}", cors.AllowAll().Handler(http_swagger.Handler(http_swagger.PersistAuthorization(true))))

	if err := srv.Start(ctx); err != nil {
		slog.Error("failed to listen http server", "err", err)
	} else {
		slog.Info("server stopped gracefully")
	}

	slog.Info("grpc server stopped gracefully")
}
