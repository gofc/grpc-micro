package server

import (
	"context"
	"github.com/google/uuid"
	"time"
)

var (
	DefaultAddress          = ":50000"
	DefaultName             = "grpc.micro.server"
	DefaultVersion          = "latest"
	DefaultId               = uuid.New().String()
	DefaultRegisterCheck    = func(context.Context) error { return nil }
	DefaultRegisterInterval = time.Second * 30
	DefaultRegisterTTL      = time.Minute
)
