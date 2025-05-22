package application

import (
	"time"

	activity_domain "github.com/sgace/backend/activity/domain"
)

func (s *ActivityService) UpdateActivity(activity_id string, activity *activity_domain.Activity) error {

	act, err := s.activityRepo.GetActivityByID(activity_id)
	if err != nil {
		return err
	}

	activity_to_postgres := activity_domain.Activity{
		ID:        activity_id,
		Name:      activity.Name,
		Type:      activity.Type,
		CreatedAt: act.CreatedAt,
		UpdatedAt: time.Now(),
		Date:      activity.Date,
		Time:      activity.Time,
	}

	return s.activityRepo.UpdateActivity(&activity_to_postgres)
}
