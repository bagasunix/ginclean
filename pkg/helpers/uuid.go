package helpers

import (
	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

func GenerateUUIDV1(logs zap.Logger) uuid.UUID {
	id, err := uuid.NewV1()
	errors.HandlerReturnedVoid(err, "uuid", "generator")
	return id
}

func GenerateUUIDV4(logs zap.Logger) uuid.UUID {
	id, err := uuid.NewV4()
	errors.HandlerReturnedVoid(err, "uuid", "generator")
	return id
}
