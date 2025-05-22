package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

func GetAllSolicitudesByUser(ctx *gin.Context) {
	var (
		uid     any
		exists  bool
		user_id string
		err     error
	)

	uid, exists = ctx.Params.Get("user_id")
	if !exists {
		ctx.AbortWithStatus(401)
		return
	}

	user_id = fmt.Sprint(uid)

	solicitude_Service := di_container.SolicitudeService()

	solicitudes, err := solicitude_Service.GetAllSolicitudesByUser(user_id)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
		}
	}

	ctx.AbortWithStatusJSON(200, gin.H{"solicitudes": solicitudes})
}
