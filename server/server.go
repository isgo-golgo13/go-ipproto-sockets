package main

import (
	"context"
	"os"

	slog "github.com/Sirupsen/logrus"
)

type ServerSvc interface {
}

type IPServer struct {
	ServerCtx context.Context

	Logger *slog.Logger
}

func NewIPServer() *IPServer {
	server := &IPServer{
		Logger: &slog.Logger{Out: os.Stderr, Formatter: &slog.TextFormatter{ForceColors: true, FullTimestamp: true}, Hooks: make(slog.LevelHooks), Level: slog.InfoLevel | slog.ErrorLevel},
	}
	server.Logger.Infof("IP Server: %s", "")
	return server
}

func init() {

}

func (s *IPServer) Loop() {

}
