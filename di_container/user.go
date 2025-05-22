package dicontainer

import (
	user_service "github.com/sgace/backend/user/application"
	user_infra "github.com/sgace/backend/user/infra"
	"github.com/sgace/database"
)

func UserService() *user_service.UserService {
	repo := user_infra.NewUserPostgresRepository(database.Database)
	return user_service.NewUserService(repo)
}
