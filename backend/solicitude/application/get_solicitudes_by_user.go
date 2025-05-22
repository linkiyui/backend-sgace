package application

import (
	solicitude_domain "github.com/sgace/backend/solicitude/domain"
)

func (s *SolicitudeService) GetAllSolicitudesByUser(user_id string) ([]solicitude_domain.Solicitude, error) {
	return s.SolicitudeRepo.GetAllSolicitudesByUser(user_id)
}
