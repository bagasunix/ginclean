package errors

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

const (
	ERR_SOMETHING_WRONG = "oops something wrong. dont worry we will fix it"
	ERR_NOT_FOUND       = "not found"
	ERR_DUPLICATE_KEY   = "duplicate key"
	ERR_ALREADY_EXISTS  = "is already exists"
	ERR_INVALID_KEY     = "invalid"
	ERR_NOT_AUTHORIZED  = "unauthorized"
)

func CustomError(err string) error {
	return errors.New(err)
}

func ErrRecordNotFound(entity string, err error) error {
	if err == nil {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprint(entity, " ", ERR_NOT_FOUND))
	}
	return ErrSomethingWrong(err)
}

func ErrDuplicateValue(entity string, err error) error {
	if err == nil {
		return err
	}
	if strings.Contains(strings.ToLower(err.Error()), ERR_DUPLICATE_KEY) {
		return errors.New(fmt.Sprint(entity, " ", ERR_ALREADY_EXISTS))
	}
	return ErrSomethingWrong(err)
}

func ErrSomethingWrong(err error) error {
	if err == nil {
		return err
	}
	return errors.New(ERR_SOMETHING_WRONG)
}

func ErrInvalidAttributes(attributes string) error {
	return errors.New(fmt.Sprint(ERR_INVALID_KEY, " ", attributes))
}
