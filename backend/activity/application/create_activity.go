package application

import (
	activity_domain "github.com/sgace/backend/activity/domain"
	"github.com/sgace/utils"
)

func (s *ActivityService) CreateActivity(activity *activity_domain.Activity) error {

	activity_ID, err := utils.GenerateUUIDv7()
	if err != nil {
		return err
	}

	activity.ID = activity_ID

	return s.activityRepo.CreateActivity(activity)
}
