package errors

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

func HandlerWithLoggerReturnedVoid(logger *zap.Logger, err error, args ...interface{}) {
	if err == nil {
		return
	}
	logger.Error(err.Error(), zap.Any("args:", args))
}
func HandlerWithOSExit(logger *zap.Logger, err error, args ...interface{}) {
	if err == nil {
		return
	}
	logger.Fatal(err.Error(), zap.Any("args:", args))
	os.Exit(1)
}

func HandlerReturnedVoid(err error, args ...interface{}) {
	if err == nil {
		return
	}
	fmt.Println("err", err)
	return
}
