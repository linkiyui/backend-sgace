package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

func DeleteMyUser(c *gin.Context) {

	var (
		user_id string
		uid     any
		err     error
		exists  bool
	)

	uid, exists = c.Params.Get("user_id")
	if !exists {
		c.AbortWithStatus(401)
	}

	user_id = fmt.Sprint(uid)

	user_service := di_container.UserService()
	err = user_service.DeleteMyUser(user_id)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			c.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		c.JSON(500, gin.H{"error": "failed to delete user"})
		return
	}

	c.AbortWithStatus(200)
}
