package application

import (
	activity_domain "github.com/sgace/backend/activity/domain"
	"github.com/sgace/errors"
)

func (s *ActivityService) GetActivityByID(id string) (*activity_domain.Activity, error) {

	activity, err := s.activityRepo.GetActivityByID(id)
	if err != nil {
		return &activity_domain.Activity{}, err
	}

	if activity == nil {
		return nil, errors.ErrNotFound
	}

	return activity, nil

}
