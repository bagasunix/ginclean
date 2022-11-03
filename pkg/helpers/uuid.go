package helpers

import (
	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/gofrs/uuid"
)

func GenerateUUIDV1() uuid.UUID {
	id, err := uuid.NewV1()
	errors.HandlerReturnedVoid(err, "uuid", "generator")
	return id
}

func GenerateUUIDV4() uuid.UUID {
	id, err := uuid.NewV4()
	errors.HandlerReturnedVoid(err, "uuid", "generator")
	return id
}
