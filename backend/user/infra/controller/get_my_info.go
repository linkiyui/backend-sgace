package controller

import (
	"fmt"

	di_container "github.com/sgace/di_container"

	"github.com/gin-gonic/gin"
	domain_error "github.com/sgace/errors"
)

func GetMyInfo(ctx *gin.Context) {

	var (
		err     error
		uid     any
		user_id string
		exists  bool
	)

	uid, exists = ctx.Params.Get("user_id")
	if !exists {
		ctx.AbortWithStatus(401)
		return
	}

	user_id = fmt.Sprint(uid)

	user_service := di_container.UserService()

	user_info, err := user_service.GetMyUserInfo(user_id)
	if err != nil {
		if de := domain_error.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
		}
	}
	info := Info{
		Id:       user_info.ID,
		Username: user_info.Username,
		Email:    user_info.Email,
		Role:     string(user_info.Role),
	}

	ctx.JSON(200, gin.H{
		"user_info": info,
	})

}

type Info struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
