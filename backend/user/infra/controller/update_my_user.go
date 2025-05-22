package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sgace/backend/user/domain"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

type UserUpdateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UpdateMyUser(c *gin.Context) {
	var (
		user_id      string
		uid          any
		err          error
		exists       bool
		user_service = di_container.UserService()
	)

	uid, exists = c.Params.Get("user_id")
	if !exists {
		c.AbortWithStatus(401)
	}

	user_id = fmt.Sprint(uid)

	var userUpdateRequest UserUpdateRequest

	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil {
		de := domain_errors.ErrbadRequest
		c.JSON(de.Code, gin.H{"error": de.Message})
		return
	}

	user := domain.User{
		Username: userUpdateRequest.Username,
		Email:    userUpdateRequest.Email,
		Password: userUpdateRequest.Password,
	}

	err = user_service.UpdateMyUser(user_id, &user)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			c.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		c.JSON(500, gin.H{"error": "failed to update user"})
		return
	}

	c.AbortWithStatus(200)

}
