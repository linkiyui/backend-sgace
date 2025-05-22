package infra

import (
	activity_domain "github.com/sgace/backend/activity/domain"
	"gorm.io/gorm"
)

type activityPostgresRepository struct {
	db *gorm.DB
}

func NewActivityPostgresRepository(db *gorm.DB) *activityPostgresRepository {
	return &activityPostgresRepository{
		db: db,
	}
}

func (r *activityPostgresRepository) CreateActivity(activity *activity_domain.Activity) error {
	if err := r.db.Table("activity").Create(activity).Error; err != nil {
		return err
	}
	return nil
}

func (r *activityPostgresRepository) UpdateActivity(activity *activity_domain.Activity) error {
	if err := r.db.Table("activity").Where("id = ?", activity.ID).Save(activity).Error; err != nil {
		return err
	}
	return nil
}

func (r *activityPostgresRepository) GetActivityByID(id string) (*activity_domain.Activity, error) {
	var activity activity_domain.Activity
	if err := r.db.Table("activity").Where("id = ?", id).First(&activity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &activity, nil
}

func (r *activityPostgresRepository) DeleteActivity(activity_id string) error {
	if err := r.db.Table("activity").Where("id = ?", activity_id).Delete(&activity_domain.Activity{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *activityPostgresRepository) GetAllActivities() ([]activity_domain.Activity, error) {
	var activities []activity_domain.Activity
	if err := r.db.Table("activity").Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}
