package main

import (
	"errors"
	"go-tools/log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger = getLogger()

func getLogger() *zap.Logger {
	// l := log.NewLumberjackLogger{"log.log", 128, 30, 7, true}
	logger := log.NewLogger(nil, log.DefaultEncoderConfig, zapcore.DebugLevel, "json", "cloud")
	return logger
}

func function() {
	log.Info("msg")
}
func main() {
	logger.Debug("ssss")
	logger.Info("ssss")
	logger.Error("ssss")
	err := errors.New("test error")
	logger.Error("ssss", zap.Any("err", err.Error()))
	log.Debug("llllll")
	function()
}
