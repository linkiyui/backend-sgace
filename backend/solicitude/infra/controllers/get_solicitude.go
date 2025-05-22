package controllers

import (
	"github.com/gin-gonic/gin"
	di_container "github.com/sgace/di_container"
)

func GetSolicitudeByID(ctx *gin.Context) {
	var (
		err error
	)

	solicitudeService := di_container.SolicitudeService()

	solicitudeID := ctx.Param("id")

	_, err = solicitudeService.GetSolicitudeByID(solicitudeID)
	if err != nil {
		ctx.AbortWithStatusJSON(404, gin.H{"error": "Solicitud not found"})
	}

	ctx.JSON(200, gin.H{"message": "Solicitude found"})
}
