package requests

import (
	"encoding/json"

	"github.com/bagasunix/ginclean/pkg/errors"
)

type BaseList struct {
	Limit   int64  `json:"limit"`
	Keyword string `json:"keyword"`
}

func (c *BaseList) ToJSON() []byte {
	j, err := json.Marshal(c)
	errors.HandlerReturnedVoid(err)
	return j
}

// BaseListBuilder Builder Object for BaseList
type BaseListBuilder struct {
	limit   int64
	keyword string
}

// NewBaseListBuilder Constructor for BaseListBuilder
func NewBaseListBuilder() *BaseListBuilder {
	o := new(BaseListBuilder)
	return o
}

// Build Method which creates BaseList
func (b *BaseListBuilder) Build() *BaseList {
	o := new(BaseList)
	o.Limit = b.limit
	o.Keyword = b.keyword
	return o
}

// SetLimit Limit Builder method to set the field limit in BaseListBuilder
func (b *BaseListBuilder) SetLimit(v int64) *BaseListBuilder {
	b.limit = v
	return b
}

// SetKeyword Keyword Builder method to set the field keyword in BaseListBuilder
func (b *BaseListBuilder) SetKeyword(v string) *BaseListBuilder {
	b.keyword = v
	return b
}
