package controller

import (
	"github.com/gin-gonic/gin"
	activity_domain "github.com/sgace/backend/activity/domain"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

func GetActivityByID(c *gin.Context) {
	var (
		err         error
		activity_id string
		activity    *activity_domain.Activity
	)

	activity_service := di_container.ActivityService()

	activity_id = c.Param("id")

	activity, err = activity_service.GetActivityByID(activity_id)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			c.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
		}
	}

	c.JSON(200, gin.H{"activity": activity})

}
