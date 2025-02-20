package main

import (
	"context"
	"fmt"
	"log/slog"
	"main/http/swagger"
	v1 "main/http/v1"
	"main/internal/disciplines"
	"main/internal/schedules"
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
}

type Tokens struct {
	Secret     string        `mapstructure:"tokenSecret"`
	AccessTTL  time.Duration `mapstructure:"accessTTL"`
	RefreshTTL time.Duration `mapstructure:"refreshTTL"`
}

const ServiceName = "schedules"

// @title           Schedules MS
// @version         1.0
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
		disciplineRepo = disciplines.NewRepo(db)
		lessonsRepo    = schedules.NewRepo(db)
	)

	// инициализация http сервера
	srv := server.New(cfg.Server.Host, cfg.Server.Port, cfg.Server.ShutdownTimeout)

	srv.Post("/api/disciplines", mw.WrapErrDetail(v1.CreateDiscipline(disciplineRepo)))
	srv.Get("/api/disciplines", mw.WrapErrDetail(v1.GetDisciplines(disciplineRepo)))
	srv.Delete("/api/disciplines/{id}", mw.WrapErrDetail(v1.DeleteDiscipline(disciplineRepo)))

	srv.Post("/api/lessons", mw.WrapErrDetail(v1.CreateLesson(lessonsRepo)))
	srv.Put("/api/lessons/{id}", mw.WrapErrDetail(v1.UpdateLesson(lessonsRepo)))
	srv.Get("/api/lessons", mw.WrapErrDetail(v1.GetLessons(lessonsRepo)))
	srv.Get("/api/lessons/{id}", mw.WrapErrDetail(v1.GetLesson(lessonsRepo)))
	srv.Delete("/api/lessons/{id}", mw.WrapErrDetail(v1.DeleteLesson(lessonsRepo)))

	swagger.SwaggerInfo.BasePath = fmt.Sprintf("/%s/", ServiceName)
	srv.Get("/swagger/{*}", cors.AllowAll().Handler(http_swagger.Handler(http_swagger.PersistAuthorization(true))))

	if err := srv.Start(ctx); err != nil {
		slog.Error("failed to listen http server", "err", err)
	} else {
		slog.Info("server stopped gracefully")
	}

	slog.Info("grpc server stopped gracefully")
}
