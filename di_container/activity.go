package dicontainer

import (
	activity_service "github.com/sgace/backend/activity/application"
	activity_infra "github.com/sgace/backend/activity/infra"
	"github.com/sgace/database"
)

func ActivityService() *activity_service.ActivityService {
	repo := activity_infra.NewActivityPostgresRepository(database.Database)
	return activity_service.NewActivityService(repo)
}
