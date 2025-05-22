package application

import (
	domain_errors "github.com/sgace/errors"
)

func (s *ActivityService) DeleteActivity(activity_id string) error {

	act, err := s.activityRepo.GetActivityByID(activity_id)
	if err != nil {
		return err
	}

	if act == nil {
		return domain_errors.ErrNotFound
	}

	return s.activityRepo.DeleteActivity(activity_id)
}
