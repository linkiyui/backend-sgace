package infra

import (
	"time"

	solicitude_domain "github.com/sgace/backend/solicitude/domain"
	"gorm.io/gorm"
)

type solicitudePostgresRepository struct {
	db *gorm.DB
}

func NewSolicitudePostgresRepository(db *gorm.DB) *solicitudePostgresRepository {
	return &solicitudePostgresRepository{
		db: db,
	}
}

func (r *solicitudePostgresRepository) CreateSolicitude(solicitude *solicitude_domain.Solicitude) error {
	err := r.db.Table("solicitude").Create(&solicitude).Error
	if err != nil {

		return err

	}

	return nil
}

func (r *solicitudePostgresRepository) GetSolicitudeByID(id string) (*solicitude_domain.Solicitude, error) {
	var solicitude solicitude_domain.Solicitude
	err := r.db.Table("solicitudes").Where("id = ?", id).First(&solicitude).Error
	if err != nil {
		return nil, err
	}
	return &solicitude, nil
}

func (r *solicitudePostgresRepository) UpdateSolicitude(solicitude *solicitude_domain.Solicitude) error {
	err := r.db.Table("solicitudes").Where("id = ?", solicitude.ID).Updates(solicitude).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *solicitudePostgresRepository) GetAllSolicitudes() ([]solicitude_domain.Solicitude, error) {
	var solicitudes []solicitude_domain.Solicitude
	err := r.db.Table("solicitudes").Find(&solicitudes).Error
	if err != nil {
		return nil, err
	}
	return solicitudes, nil
}

func (r *solicitudePostgresRepository) GetAllSolicitudesByUser(user_id string) ([]solicitude_domain.Solicitude, error) {
	var solicitudes []solicitude_domain.Solicitude
	err := r.db.Table("solicitudes").Where("user_id = ?", user_id).Find(&solicitudes).Error
	if err != nil {
		return nil, err
	}
	return solicitudes, nil
}

func (r *solicitudePostgresRepository) DeleteSolicitude(id string) error {
	err := r.db.Table("solicitudes").Where("id = ?", id).Delete(&solicitude_domain.Solicitude{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *solicitudePostgresRepository) GetAcceptedSolicitudes(id string) ([]solicitude_domain.Solicitude, error) {
	var solicitudes []solicitude_domain.Solicitude
	err := r.db.Table("solicitudes").Where("status = ?", solicitude_domain.Accepted, "user_id = ?", id).Find(&solicitudes).Error
	if err != nil {
		return nil, err
	}
	return solicitudes, nil
}

func (r *solicitudePostgresRepository) GetRefusedSolicitudes(id string) ([]solicitude_domain.Solicitude, error) {
	var solicitudes []solicitude_domain.Solicitude
	err := r.db.Table("solicitudes").Where("status = ?", solicitude_domain.Rejected, "user_id = ?", id).Find(&solicitudes).Error
	if err != nil {
		return nil, err
	}
	return solicitudes, nil
}

func (r *solicitudePostgresRepository) GetComppletedSolicitudes(id string) ([]solicitude_domain.Solicitude, error) {
	var solicitudes []solicitude_domain.Solicitude
	err := r.db.Table("solicitudes").Where("date <= ?", time.Now().Format("2006-01-02"), "user_id = ?", id).Find(&solicitudes).Error
	if err != nil {
		return nil, err
	}
	return solicitudes, nil
}
