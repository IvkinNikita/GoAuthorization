package grpc

import (
	"log/slog" 

	"google.golang.org/grpc"
)

type App struct {
	log         *slog.Logger
	gRPCSerever *grpc.Server
	port        string
}

func New