package dicontainer

import (
	auth_service "github.com/sgace/backend/auth/application"
)

func AuthService() *auth_service.AuthService {
	userSrv := UserService()
	return auth_service.NewAuthService(userSrv)

}
