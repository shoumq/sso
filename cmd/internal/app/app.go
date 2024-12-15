package app

import (
	"log/slog"
	grpcapp "sso/cmd/internal/app/grpc"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTl time.Duration,
) *App {
	// TODO: Storage

	// TODO: Init auth service

	grpcApp := grpcapp.New(log, grpcPort)
	return &App{GRPCServer: grpcApp}
}
