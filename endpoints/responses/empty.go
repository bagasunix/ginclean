package responses

import (
	"encoding/json"

	"github.com/bagasunix/ginclean/pkg/errors"
)

type Empty struct{}

func (a *Empty) ToJSON() []byte {
	j, err := json.Marshal(a)
	errors.HandlerReturnedVoid(err)
	return j
}
