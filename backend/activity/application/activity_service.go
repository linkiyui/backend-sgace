package application

import (
	activity_domain "github.com/sgace/backend/activity/domain"
)

type IActivityService interface {
	CreateActivity(activity *activity_domain.Activity) error
	UpdateActivity(activity *activity_domain.Activity) error
	DeleteActivity(activity_id string)
	GetAllActivities() ([]activity_domain.Activity, error)
	GetActivityByID(id string) (*activity_domain.Activity, error)
	GetTotalActivities() (int64, error)
}

type ActivityService struct {
	activityRepo activity_domain.ISolicitudeRepository
}

func NewActivityService(activityRepo activity_domain.ISolicitudeRepository) *ActivityService {
	return &ActivityService{
		activityRepo: activityRepo,
	}
}

func (s *ActivityService) GetTotalActivities() (int64, error) {
	var (
		err        error
		activities []activity_domain.Activity
	)

	activities, err = s.activityRepo.GetAllActivities()
	if err != nil {
		return 0, err
	}

	return int64(len(activities)), nil
}
