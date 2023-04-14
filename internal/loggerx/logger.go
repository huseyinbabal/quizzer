package loggerx

import "go.uber.org/zap"

func New() (*zap.Logger, error) {
	return zap.NewProduction()
}
