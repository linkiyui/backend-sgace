package application

import (
	domain_errors "github.com/sgace/errors"
)

func (s *SolicitudeService) DeleteSolicitude(id string) error {

	solicitude, err := s.SolicitudeRepo.GetSolicitudeByID(id)
	if err != nil {
		return err
	}

	if solicitude == nil {
		return domain_errors.ErrNotFound
	}

	err = s.SolicitudeRepo.DeleteSolicitude(id)
	if err != nil {
		return err
	}

	return nil

}
