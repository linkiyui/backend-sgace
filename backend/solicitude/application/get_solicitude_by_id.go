package application

import (
	"github.com/sgace/backend/solicitude/domain"
	domain_errors "github.com/sgace/errors"
)

func (s *SolicitudeService) GetSolicitudeByID(id string) (*domain.Solicitude, error) {

	solicitud, err := s.SolicitudeRepo.GetSolicitudeByID(id)
	if err != nil {
		return nil, err
	}

	if solicitud == nil {
		return nil, domain_errors.ErrNotFound
	}

	return solicitud, nil

}
