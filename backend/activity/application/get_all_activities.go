package application

import (
	activity_domain "github.com/sgace/backend/activity/domain"
)

func (s *ActivityService) GetAllActivities() ([]activity_domain.Activity, error) {

	var (
		err        error
		activities []activity_domain.Activity
	)

	activities, err = s.activityRepo.GetAllActivities()
	if err != nil {

		return nil, err
	}

	return activities, nil

}
