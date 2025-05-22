package controller

import (
	"github.com/gin-gonic/gin"
	activity_domain "github.com/sgace/backend/activity/domain"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

func GetAllActivities(c *gin.Context) {

	var (
		err        error
		activities []activity_domain.Activity
	)

	activity_service := di_container.ActivityService()

	activities, err = activity_service.GetAllActivities()
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			c.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
		}
	}

	c.JSON(200, gin.H{"activities": activities})

}
