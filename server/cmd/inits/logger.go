package inits

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lock = &sync.Mutex{}
var Loggers *zap.Logger

func InitLogger() *zap.Logger {
	var core zapcore.Core

	core = zapcore.NewTee(
		zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)

	Loggers = zap.New(core, zap.AddCaller())

	return Loggers
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
