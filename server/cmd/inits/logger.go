package inits

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"sync"

	envs "github.com/bagasunix/ginclean/pkg/env"
	"github.com/bagasunix/ginclean/pkg/helpers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var lock = &sync.Mutex{}
var Loggers *zap.Logger

func InitLogger(config *envs.Configs) *zap.Logger {
	fmt.Println(config.LogPath)
	writerSyncer := getLogWriter(config)
	var core zapcore.Core

	if os.Getenv("ENV") == "dev" {
		core = zapcore.NewTee(
			zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(getFileEncoder(), writerSyncer, zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(getFileEncoder(), writerSyncer, zapcore.DebugLevel)
	}

	Loggers = zap.New(core, zap.AddCaller())

	return Loggers
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getFileEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.StacktraceKey = zapcore.DPanicLevel.String()
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(config *envs.Configs) zapcore.WriteSyncer {

	logFilePath := config.LogPath
	logFileName := config.LogName + "-" + helpers.DateFormatEN() + ".log"
	logFileMaxSize, _ := strconv.Atoi(config.LogMaxSize)
	logFileMaxBackups, _ := strconv.Atoi(config.LogMaxBackup)
	logFileMaxAge, _ := strconv.Atoi(config.LogMaxAge)
	logFile := path.Join(logFilePath, logFileName)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    logFileMaxSize,
		MaxBackups: logFileMaxBackups,
		MaxAge:     logFileMaxAge,
		Compress:   true,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
