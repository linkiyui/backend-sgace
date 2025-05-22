package application

import solicitude_domain "github.com/sgace/backend/solicitude/domain"

type ISolicitudeService interface {
	CreateSolicitude(solicitude *solicitude_domain.Solicitude) error
	AcceptSolicitude(solicitude_id string) error
	SaveSolicitude(solicitude *solicitude_domain.Solicitude) error
	RefuseSolicitude(solicitude_id string) error
	GetAllSolicitudes() ([]solicitude_domain.Solicitude, error)
	GetSolicitudeByID(id string) (*solicitude_domain.Solicitude, error)
	DeleteSolicitude(id string) error
	GetAllSolicitudesByUser(user_id string) ([]solicitude_domain.Solicitude, error)
	UpdateSolicitude(*solicitude_domain.Solicitude) error
	GetTotalSolicitudes(id string) (int64, error)
	GetAcceptedSolicitudes(id string) ([]solicitude_domain.Solicitude, error)
	GetRefusedSolicitudes(id string) (int64, error)
	GetCompletedSolicitudes(id string) (int64, error)
}

type SolicitudeService struct {
	SolicitudeRepo solicitude_domain.ISolicitudeRepository
}

func NewSolicitudeService(solicitudeRepo solicitude_domain.ISolicitudeRepository) *SolicitudeService {
	return &SolicitudeService{
		SolicitudeRepo: solicitudeRepo,
	}
}

func (s *SolicitudeService) GetTotalSolicitudes(id string) (int64, error) {
	var (
		err         error
		solicitudes []solicitude_domain.Solicitude
	)

	solicitudes, err = s.SolicitudeRepo.GetAllSolicitudesByUser(id)
	if err != nil {
		return 0, err
	}

	return int64(len(solicitudes)), nil
}

func (s *SolicitudeService) GetRefusedSolicitudes(id string) (int64, error) {
	var (
		err         error
		solicitudes []solicitude_domain.Solicitude
	)

	solicitudes, err = s.SolicitudeRepo.GetRefusedSolicitudes(id)
	if err != nil {
		return 0, err
	}

	return int64(len(solicitudes)), nil
}

func (s *SolicitudeService) GetCompletedSolicitudes(id string) (int64, error) {
	var (
		err         error
		solicitudes []solicitude_domain.Solicitude
	)

	solicitudes, err = s.SolicitudeRepo.GetComppletedSolicitudes(id)
	if err != nil {
		return 0, err
	}

	return int64(len(solicitudes)), nil
}
