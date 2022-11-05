package endpoints

import (
	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func decodeByEntityIdEndpoint(g *gin.Context) (request interface{}, err error) {
	vars := g.Param("id")
	if vars == "" {
		return nil, errors.ErrInvalidAttributes("id")
	}
	uuidId, err := uuid.FromString(vars)
	requestBuilder := requests.NewEntityIdBuilder()
	requestBuilder.SetId(uuidId)
	return requestBuilder.Build(), err
}
