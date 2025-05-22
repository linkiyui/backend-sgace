package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sgace/backend/activity/domain"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

func UpdateActivity(ctx *gin.Context) {

	var (
		err         error
		activity_id string
	)

	activity_id = ctx.Param("id")
	if activity_id == "" {
		de := domain_errors.ErrbadRequest
		ctx.JSON(de.Code, gin.H{"error": de.Message})
		return
	}

	var req CreateActivityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		de := domain_errors.ErrbadRequest
		ctx.JSON(de.Code, gin.H{"error": de.Message})
		return
	}

	activity_service := di_container.ActivityService()

	activity := domain.Activity{
		Name:      req.Name,
		Type:      req.Type,
		Date:      req.Date,
		Time:      req.Tiime,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = activity_service.UpdateActivity(activity_id, &activity)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		ctx.JSON(500, gin.H{"error": "failed to update activity"})
		return
	}

	ctx.AbortWithStatus(200)

}
