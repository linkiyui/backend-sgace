package controllers

import (
	"github.com/gin-gonic/gin"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

func AcceptSolicitude(ctx *gin.Context) {
	var (
		err error
	)

	solicitudeService := di_container.SolicitudeService()

	solicitudeID := ctx.Param("id")

	err = solicitudeService.AcceptSolicitude(solicitudeID)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
		}
	}

	ctx.AbortWithStatus(200)
}
