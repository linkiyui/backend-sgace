package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

func DeleteSolicitudeById(ctx *gin.Context) {
	var (
		err error
	)

	uid, exists := ctx.Params.Get("id")

	if !exists {
		ctx.AbortWithStatus(401)
		return
	}

	id := fmt.Sprint(uid)

	solicitude_Service := di_container.SolicitudeService()

	err = solicitude_Service.DeleteSolicitude(id)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		ctx.AbortWithStatusJSON(404, gin.H{"error": "Solicitude not found"})
		return
	}

	ctx.AbortWithStatus(200)
}
