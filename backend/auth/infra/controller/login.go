package controller

import (
	"github.com/gin-gonic/gin"
	auth_token "github.com/sgace/backend/auth/auth_token"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	auth_service := di_container.AuthService()

	user, err := auth_service.Login(req.Username, req.Password)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		ctx.JSON(500, gin.H{"error": "failed to login"})
		return
	}

	token, err := auth_token.GenerateLoginToken(user.ID, string(user.Role))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}
	refresh_token, err := auth_token.GenerateRefreshToken(user.ID, string(user.Role))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to generate refresh token"})
		return
	}
	ctx.JSON(200, gin.H{
		"token":         token,
		"refresh_token": refresh_token,
		"user":          user,
	})

	// TODO: implement login logic
}
