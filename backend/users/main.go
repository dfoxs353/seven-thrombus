package main

import (
	"context"
	"fmt"
	"log/slog"
	"main/grpc"
	"main/http/swagger"
	v1 "main/http/v1"
	"main/internal/jwt"
	"main/internal/middleware"
	"main/internal/users"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/cors"
	http_swagger "github.com/swaggo/http-swagger"
	accounts_proto "gitlab.com/volgaIt/grpc-proto/accounts"
	config "gitlab.com/volgaIt/packages/config"
	mw "gitlab.com/volgaIt/packages/middleware"
	"gitlab.com/volgaIt/packages/postgres"
	server "gitlab.com/volgaIt/packages/server"
	googlegrpc "google.golang.org/grpc"
)

type cfg struct {
	Postgres     config.Postgres   `mapstructure:"postgres"`
	Server       config.Server     `mapstructure:"server"`
	GrpcServer   config.GrpcServer `mapstructure:"grpcServer"`
	Tokens       Tokens            `mapstructure:"tokens"`
	Cost         int               `mapstructure:"cost"`
	DefaultRoles []string          `mapstructure:"defaultRoles"`
}

type Tokens struct {
	Secret     string        `mapstructure:"tokenSecret"`
	AccessTTL  time.Duration `mapstructure:"accessTTL"`
	RefreshTTL time.Duration `mapstructure:"refreshTTL"`
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
		userRepo    = users.NewRepo(db)
		jwtManager  = jwt.New(cfg.Tokens.Secret, cfg.Tokens.AccessTTL, cfg.Tokens.RefreshTTL)
		userService = users.NewService(userRepo, cfg.Cost, jwtManager, db)

		// user handlers
		signUp        = v1.SignUp(userService, cfg.DefaultRoles)
		signIn        = v1.SignIn(userService)
		signOut       = v1.SignOut(userRepo)
		validate      = v1.Validate(jwtManager)
		refresh       = v1.Refresh(userService)
		getProfile    = v1.Profile(userRepo)
		updateProfile = v1.UpdateProfile(userService)

		// admin handlers
		getUsers   = v1.GetUsers(userRepo)
		createUser = v1.CreateUser(userService)
		updateUser = v1.UpdateUser(userService)
		deleteUser = v1.DeleteUserSoft(userRepo)
	)

	// signup default users
	_, err = userService.SignUp(ctx, "admin", "admin", "Админ", "Админов", []users.Role{users.Admin})
	if err != nil {
		slog.WarnContext(ctx, "", "err", err)
	}

	_, err = userService.SignUp(ctx, "teacher", "teacher", "Преподаватель", "Автоматов", []users.Role{users.Teacher})
	if err != nil {
		slog.WarnContext(ctx, "", "err", err)
	}

	_, err = userService.SignUp(ctx, "student", "student", "Студент", "Отличников", []users.Role{users.Student})
	if err != nil {
		slog.WarnContext(ctx, "", "err", err)
	}

	// инициализация http сервера
	srv := server.New(cfg.Server.Host, cfg.Server.Port, cfg.Server.ShutdownTimeout)

	srv.Post("/api/signup", mw.WrapErrDetail(signUp))
	srv.Post("/api/signin", mw.WrapErrDetail(signIn))
	srv.Put("/api/signout", middleware.WrapAuth(mw.WrapErrDetail(signOut), jwtManager, nil))
	srv.Get("/api/validate", mw.WrapErrDetail(validate))
	srv.Post("/api/refresh", mw.WrapErrDetail(refresh))

	srv.Get("/api/accounts/me", middleware.WrapAuth(mw.WrapErrDetail(getProfile), jwtManager, nil))
	srv.Put("/api/accounts/update", middleware.WrapAuth(mw.WrapErrDetail(updateProfile), jwtManager, nil))

	// admin routes
	srv.Get("/api/accounts", middleware.WrapAuth(mw.WrapErrDetail(getUsers), jwtManager, []users.Role{users.Admin}))
	srv.Post("/api/accounts", middleware.WrapAuth(mw.WrapErrDetail(createUser), jwtManager, []users.Role{users.Admin}))
	srv.Put("/api/accounts/{id}", middleware.WrapAuth(mw.WrapErrDetail(updateUser), jwtManager, []users.Role{users.Admin}))
	srv.Delete("/api/accounts/{id}", middleware.WrapAuth(mw.WrapErrDetail(deleteUser), jwtManager, []users.Role{users.Admin}))

	swagger.SwaggerInfo.BasePath = fmt.Sprintf("/%s/", ServiceName)
	srv.Get("/swagger/{*}", cors.AllowAll().Handler(http_swagger.Handler(http_swagger.PersistAuthorization(true))))

	// инициализация grpc сервера
	grpcSrv := googlegrpc.NewServer()
	accounts_proto.RegisterAccountsServer(grpcSrv, grpc.New(jwtManager, userRepo))
	go func() {
		addr := fmt.Sprintf("%s:%d", cfg.GrpcServer.Host, cfg.GrpcServer.Port)
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			slog.ErrorContext(ctx, "failed to make listener", "err", err)
			cancel()
			return
		}
		defer lis.Close()

		slog.Info("grpc server successfully started", "addr", addr)

		if err := grpcSrv.Serve(lis); err != nil {
			slog.ErrorContext(ctx, "failed to serve grpc server", "err", err)
			cancel()
			return
		}
	}()

	if err := srv.Start(ctx); err != nil {
		slog.Error("failed to listen http server", "err", err)
	} else {
		slog.Info("server stopped gracefully")
	}

	grpcSrv.GracefulStop()
	slog.Info("grpc server stopped gracefully")
}
