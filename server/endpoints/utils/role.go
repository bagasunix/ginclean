package utils

import (
	"github.com/bagasunix/ginclean/pkg/errors"
	"github.com/bagasunix/ginclean/server/endpoints/requests"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func DecodeByUpdateRoleEndpoint(g *gin.Context) (request interface{}, err error) {
	vars := g.Param("id")
	if vars == "" {
		return nil, errors.ErrInvalidAttributes("id")
	}
	var uuidId uuid.UUID
	if uuidId, err = uuid.FromString(vars); err != nil {
		return nil, errors.ErrInvalidAttributes("id")
	}

	var req requests.UpdateRole

	err = g.Bind(&req)
	req.Id = uuidId
	return &req, err
}
