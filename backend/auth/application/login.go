package application

import (
	user_domain "github.com/sgace/backend/user/domain"
	domain_errors "github.com/sgace/errors"
	"github.com/sgace/utils"
)

func (s *AuthService) Login(username, password string) (user_domain.User, error) {
	// Check if the user exists
	user, err := s.user_serv.GetUserByUsername(username)
	if err != nil {
		return user_domain.User{}, domain_errors.ErrUserNotFound
	}
	// Check if the password is correct
	hashed_password := user.Password
	if err := utils.CheckPasswordHashFunc(hashed_password, password); err == false {
		return user_domain.User{}, domain_errors.ErrInvalidPassword
	}

	return *user, nil

}
