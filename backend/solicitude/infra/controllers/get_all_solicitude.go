package controllers

import (
	"github.com/gin-gonic/gin"
	di_container "github.com/sgace/di_container"
)

func GetAllSolicitudes(ctx *gin.Context) {
	var (
		err error
	)

	solicitudeService := di_container.SolicitudeService()

	solicitudes, err := solicitudeService.GetAllSolicitudes()
	if err != nil {
		ctx.AbortWithStatusJSON(404, gin.H{"error": "Solicitudes not found"})
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"solicitudes": solicitudes,
	})
}
