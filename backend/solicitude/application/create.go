package application

import (
	solicitude_domain "github.com/sgace/backend/solicitude/domain"
	"github.com/sgace/utils"
)

func (s *SolicitudeService) CreateSolicitude(solicitude *solicitude_domain.Solicitude) error {

	id, err := utils.GenerateUUIDv7()
	if err != nil {
		return err
	}

	solicitude_to_postgres := solicitude_domain.Solicitude{
		ID:         id,
		ActivityID: solicitude.ActivityID,
		UserID:     solicitude.UserID,
		Status:     solicitude_domain.Accepted,
		Group:      solicitude.Group,
		Faculty:    solicitude.Faculty,
		Grade:      solicitude.Grade,
		CreatedAt:  solicitude.CreatedAt,
		UpdatedAt:  solicitude.UpdatedAt,
	}

	err = s.SolicitudeRepo.CreateSolicitude(&solicitude_to_postgres)
	if err != nil {
		return err
	}

	return nil
}
