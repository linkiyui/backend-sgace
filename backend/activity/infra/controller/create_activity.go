package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sgace/backend/activity/domain"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

type CreateActivityRequest struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Date  string `json:"date"`
	Tiime string `json:"time"`
}

func CreateActivity(ctx *gin.Context) {

	var req CreateActivityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		de := domain_errors.ErrbadRequest
		ctx.JSON(de.Code, gin.H{"error": de.Message})
		return
	}

	activity := domain.Activity{
		Name:      req.Name,
		Type:      req.Type,
		Date:      req.Date,
		Time:      req.Tiime,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	activity_service := di_container.ActivityService()
	err := activity_service.CreateActivity(&activity)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		ctx.JSON(500, gin.H{"error": "failed to create activity"})
		return
	}

	ctx.AbortWithStatus(200)

}
