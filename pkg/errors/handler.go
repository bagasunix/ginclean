package errors

import (
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
)

func HandlerWithOSExit(err error, args ...interface{}) {
	if err == nil {
		return
	}
	log.Fatal(err.Error(), zap.Any("args:", args))
	os.Exit(1)
}

func HandlerReturnedVoid(err error, args ...interface{}) {
	if err == nil {
		return
	}
	fmt.Println("err", err)
	return
}
