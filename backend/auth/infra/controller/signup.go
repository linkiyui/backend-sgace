package controller

import (
	"github.com/gin-gonic/gin"
	auth_token "github.com/sgace/backend/auth/auth_token"
	user_domain "github.com/sgace/backend/user/domain"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
	"github.com/sgace/utils"
)

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func Signup(ctx *gin.Context) {
	var req SignupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := utils.GenerateUUIDv7()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to generate id"})
		return
	}

	// fmt.Println(id)

	user := &user_domain.User{
		ID:       id,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     user_domain.Role(req.Role),
	}

	auth_service := di_container.AuthService()

	err = auth_service.SignUp(user)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		ctx.JSON(500, gin.H{"error": "failed to create user"})
		return
	}

	// fmt.Println(user)

	token, err := auth_token.GenerateLoginToken(user.ID, string(user.Role))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}
	refres_token, err := auth_token.GenerateRefreshToken(user.ID, string(user.Role))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to generate refresh token"})
		return
	}

	ctx.JSON(200, gin.H{
		"token":         token,
		"refresh_token": refres_token,
		"ID":            user.ID,
	})

}
