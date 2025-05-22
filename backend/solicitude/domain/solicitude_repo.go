package domain

type ISolicitudeRepository interface {
	GetSolicitudeByID(id string) (*Solicitude, error)
	UpdateSolicitude(solicitude *Solicitude) error
	GetAllSolicitudes() ([]Solicitude, error)
	CreateSolicitude(*Solicitude) error
	GetAllSolicitudesByUser(user_id string) ([]Solicitude, error)
	DeleteSolicitude(id string) error
	GetAcceptedSolicitudes(id string) ([]Solicitude, error)
	GetRefusedSolicitudes(id string) ([]Solicitude, error)
	GetComppletedSolicitudes(id string) ([]Solicitude, error)
}
