package models

import (
	"go.uber.org/zap"
	"time"
)

type Logger interface {
	Info(*HttpRequest)
	Error(*HttpRequest)
	Debug(*DbQuery)
}

type DbQuery struct {
	Message     string
	QueryString string
	XRQV        string
	XRQK        string
}

type HttpRequest struct {
	Status  int
	Message string
	XRQV    string
	XRQK    string
	Method  string
	Path    string
	Latency time.Duration
}

type LOG struct {
	*zap.Logger
}

func NewLogger() (*LOG, error) {
	logger, err := zap.NewDevelopment()
	return &LOG{logger}, err
}

func (log *LOG) Info(data *HttpRequest) {
	log.Logger.Info(data.Message,
		zap.String(data.XRQK, data.XRQV),
		zap.String("method", data.Method),
		zap.String("url", data.Path),
		zap.Duration("latency", data.Latency))
}

func (log *LOG) Error(data *HttpRequest) {
	log.Logger.Error(data.Message,
		zap.String(data.XRQK, data.XRQV),
		zap.String("method", data.Method),
		zap.String("url", data.Path),
		zap.Duration("latency", data.Latency))
}

func (log *LOG) Debug(data *DbQuery) {
	log.Logger.Debug(data.Message,
		zap.String(data.XRQK, data.XRQV),
		zap.String("query", data.QueryString))
}
