package responses

import (
	"encoding/json"

	"github.com/bagasunix/ginclean/pkg/errors"
)

type ListMultiple[T any, F any] struct {
	Data ListMultiples[T, F] `json:"data"`
}

type ListMultiples[T any, F any] struct {
	Success []T `json:"success"`
	Failed  []F `json:"failed"`
}

func (a *ListMultiple[T, F]) ToJSON() []byte {
	j, err := json.Marshal(a)
	errors.HandlerReturnedVoid(err)
	return j
}
