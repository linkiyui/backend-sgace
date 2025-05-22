package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sgace/backend/solicitude/domain"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

type Solicitude struct {
	ActivityID string `json:"activity_id"`
	Group      string `json:"group"`
	Faculty    string `json:"faculty"`
	Grade      string `json:"grade"`
}

func CreateSolicitude(ctx *gin.Context) {
	var (
		uid     any
		exists  bool
		user_id string
		err     error
	)

	uid, exists = ctx.Params.Get("user_id")
	if !exists {
		ctx.AbortWithStatus(401)
		return
	}

	user_id = fmt.Sprint(uid)

	var req Solicitude

	solicitude_Service := di_container.SolicitudeService()
	activity_service := di_container.ActivityService()

	err = ctx.BindJSON(&req)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
		}
	}

	_, err = activity_service.GetActivityByID(req.ActivityID)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		ctx.AbortWithStatusJSON(404, gin.H{"error": "Activity not found"})
		return
	}

	solicitude := domain.Solicitude{

		ActivityID: req.ActivityID,
		UserID:     user_id,
		Status:     domain.Accepted,
		Group:      req.Group,
		Faculty:    domain.Faculty(req.Faculty),
		Grade:      req.Grade,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err = solicitude_Service.CreateSolicitude(&solicitude)
	if err != nil {
		fmt.Println(err)
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
		}
	}

	ctx.AbortWithStatus(200)
}
