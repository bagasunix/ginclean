package responses

import (
	"encoding/json"

	"github.com/bagasunix/ginclean/pkg/errors"
)

type ListMultiple[T any, F any] struct {
	Data ListMultiples[T, F] `json:"data"`
}

type ListMultipleBuilder[T any, F any] struct {
	data ListMultiples[T, F]
}

func NewListMultipleBuilder[T any, F any]() *ListMultipleBuilder[T, F] {
	a := new(ListMultipleBuilder[T, F])
	return a
}

func (a *ListMultipleBuilder[T, F]) Build() *ListMultiple[T, F] {
	b := new(ListMultiple[T, F])
	b.Data.Failed = a.data.Failed
	b.Data.Success = a.data.Success
	return b
}

func (a *ListMultipleBuilder[T, F]) SetDataMulti(dataSuccess []T, dataFailed []F) *ListMultipleBuilder[T, F] {
	a.data.Success = dataSuccess
	a.data.Failed = dataFailed
	return a
}

func (a *ListMultiple[T, F]) ToJSON() []byte {
	j, err := json.Marshal(a)
	errors.HandlerReturnedVoid(err)
	return j
}

type ListMultiples[T any, F any] struct {
	Success []T `json:"success"`
	Failed  []F `json:"failed"`
}

type ListMultiplesBuilder[T any, F any] struct {
	success []T
	failed  []F
}

func NewListMultiplesBuilder[T any, F any]() *ListMultiplesBuilder[T, F] {
	a := new(ListMultiplesBuilder[T, F])
	return a
}

func (a *ListMultiplesBuilder[T, F]) Build() *ListMultiples[T, F] {
	b := new(ListMultiples[T, F])
	b.Failed = a.failed
	b.Success = a.success
	return b
}

func (a *ListMultiples[T, F]) ToJSON() []byte {
	j, err := json.Marshal(a)
	errors.HandlerReturnedVoid(err)
	return j
}
