package utils

import "go.uber.org/zap"

var Logger *zap.Logger

func SetupLogger() {
	Logger, _ = zap.NewProduction()
}
