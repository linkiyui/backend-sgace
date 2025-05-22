package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	solicitude_domain "github.com/sgace/backend/solicitude/domain"
	di_container "github.com/sgace/di_container"
	domain_errors "github.com/sgace/errors"
)

type UpdateSolicitude struct {
	ActivityID string `json:"activity_id"`
	Group      string `json:"group"`
	Faculty    string `json:"faculty"`
	Grade      string `json:"grade"`
}

func UpdateSolicitudeById(ctx *gin.Context) {

	var (
		err error
	)

	var req UpdateSolicitude
	uid, exists := ctx.Params.Get("id")

	if !exists {
		ctx.AbortWithStatus(401)
		return
	}

	id := fmt.Sprint(uid)

	err = ctx.BindJSON(&req)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
		}
	}

	solicitude_Service := di_container.SolicitudeService()

	solicitude, err := solicitude_Service.GetSolicitudeByID(id)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
		}
	}

	solicitude.ActivityID = req.ActivityID
	solicitude.Group = req.Group
	solicitude.Faculty = solicitude_domain.Faculty(req.Faculty)
	solicitude.Grade = req.Grade

	err = solicitude_Service.UpdateSolicitude(solicitude)
	if err != nil {
		if de := domain_errors.IsDomainError(err); de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		ctx.AbortWithStatusJSON(404, gin.H{"error": "Solicitude not found"})
		return
	}

	ctx.AbortWithStatus(200)

}
