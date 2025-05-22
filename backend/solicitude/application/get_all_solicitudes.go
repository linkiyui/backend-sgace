package application

import (
	"github.com/sgace/backend/solicitude/domain"
)

func (s *SolicitudeService) GetAllSolicitudes() ([]domain.Solicitude, error) {

	solicitudes, err := s.SolicitudeRepo.GetAllSolicitudes()
	if err != nil {
		return nil, err
	}

	var solicitudes_domain []domain.Solicitude

	for _, solicitud := range solicitudes {
		solicitudes_domain = append(solicitudes_domain, domain.Solicitude{
			ID:         solicitud.ID,
			ActivityID: solicitud.ActivityID,
			UserID:     solicitud.UserID,
			Status:     solicitud.Status,
			Group:      solicitud.Group,
			Faculty:    solicitud.Faculty,
			Grade:      solicitud.Grade,
			CreatedAt:  solicitud.CreatedAt,
			UpdatedAt:  solicitud.UpdatedAt,
		})
	}

	return solicitudes_domain, nil

}
