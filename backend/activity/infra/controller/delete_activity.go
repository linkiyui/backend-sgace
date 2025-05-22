package controller

import (
	"github.com/gin-gonic/gin"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

func DeleteActivity(c *gin.Context) {

	var (
		err         error
		activity_id string
	)

	activity_id = c.Param("id")
	if activity_id == "" {
		de := domain_errors.ErrbadRequest
		c.JSON(de.Code, gin.H{"error": de.Message})
		return
	}

	activity_service := di_container.ActivityService()
	err = activity_service.DeleteActivity(activity_id)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			c.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		c.JSON(500, gin.H{"error": "failed to delete activity"})
		return
	}

	c.AbortWithStatus(200)

}
