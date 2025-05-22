package application

import (
	"time"

	"github.com/sgace/backend/solicitude/domain"
	domain_errors "github.com/sgace/errors"
)

func (s *SolicitudeService) RefuseSolicitude(solicitude_id string) error {

	solicitud, err := s.SolicitudeRepo.GetSolicitudeByID(solicitude_id)
	if err != nil {
		return err
	}

	if solicitud == nil {
		return domain_errors.ErrNotFound
	}

	refused_solicitude := domain.Solicitude{
		ID:         solicitude_id,
		ActivityID: solicitud.ActivityID,
		UserID:     solicitud.UserID,
		Status:     domain.Rejected,
		Group_id:   solicitud.Group_id,
		Faculty:    solicitud.Faculty,
		Grade:      solicitud.Grade,
		CreatedAt:  solicitud.CreatedAt,
		UpdatedAt:  time.Now(),
	}

	if err := s.SolicitudeRepo.UpdateSolicitude(&refused_solicitude); err != nil {
		return err
	}

	return nil

}
