package application

import "github.com/sgace/backend/solicitude/domain"

func (s *SolicitudeService) UpdateSolicitude(solicitude *domain.Solicitude) error {
	return s.SolicitudeRepo.UpdateSolicitude(solicitude)
}
