package application

import (
	"time"

	"github.com/sgace/backend/solicitude/domain"
	domain_errors "github.com/sgace/errors"
)

func (s *SolicitudeService) AcceptSolicitude(solicitude_id string) error {

	solicitud, err := s.SolicitudeRepo.GetSolicitudeByID(solicitude_id)
	if err != nil {
		return err
	}

	if solicitud == nil {
		return domain_errors.ErrNotFound
	}

	accepted_solicitude := domain.Solicitude{
		ID:         solicitude_id,
		ActivityID: solicitud.ActivityID,
		UserID:     solicitud.UserID,
		Status:     domain.Accepted,
		Group_id:   solicitud.Group_id,
		Faculty:    solicitud.Faculty,
		Grade:      solicitud.Grade,
		CreatedAt:  solicitud.CreatedAt,
		UpdatedAt:  time.Now(),
	}

	err = s.SolicitudeRepo.UpdateSolicitude(&accepted_solicitude)
	if err != nil {
		return err
	}

	return nil

}
