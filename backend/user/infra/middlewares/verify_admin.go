package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	dicontainer "github.com/sgace/di_container"
)

func VerifyAdmin(ctx *gin.Context) {
	var (
		uid          any
		exists       bool
		user_id      string
		err          error
		user_service = dicontainer.UserService()
	)

	uid, exists = ctx.Params.Get("user_id")
	if !exists {
		ctx.AbortWithStatus(401)
		return
	}

	user_id = fmt.Sprint(uid)

	user, err := user_service.GetUserByID(user_id)
	if err != nil {
		ctx.AbortWithStatus(401)
		return
	}

	if user.Role != "admin" {
		ctx.AbortWithStatus(403)
		return
	}

	ctx.Next()

}
