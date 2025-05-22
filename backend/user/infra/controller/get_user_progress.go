package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	user_domain "github.com/sgace/backend/user/domain"
	di_container "github.com/sgace/di_container"
)

func GetUserProgress(ctx *gin.Context) {
	var (
		err error
	)

	uid, exists := ctx.Params.Get("id")

	if !exists {
		ctx.AbortWithStatus(401)
		return
	}

	id := fmt.Sprint(uid)

	activity_service := di_container.ActivityService()
	solicitude_sercvice := di_container.SolicitudeService()

	total_activities, err := activity_service.GetTotalActivities()
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	total_solicitudes, err := solicitude_sercvice.GetTotalSolicitudes(id)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	accepted_solicitudes, err := solicitude_sercvice.GetAcceptedSolicitudes(id)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	refused_solicitudes, err := solicitude_sercvice.GetRefusedSolicitudes(id)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}
	completed, err := solicitude_sercvice.GetCompletedSolicitudes(id)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	progress := user_domain.UserProgress{
		ToatlActivities:     total_activities,
		TotalSolicitudes:    total_solicitudes,
		AcceptedSolicitudes: accepted_solicitudes,
		RefuseSolicitudes:   refused_solicitudes,
		Completed:           completed,
	}

	ctx.JSON(200, progress)
}
