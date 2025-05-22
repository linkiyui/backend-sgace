package application

import (
	user_domain "github.com/sgace/backend/user/domain"
	domain_errors "github.com/sgace/errors"
)

func (s *AuthService) SignUp(user *user_domain.User) error {

	// Check if user already exists
	exists, err := s.user_serv.ExistsUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return domain_errors.ErrEmailAlreadyExists
	}

	exists, err = s.user_serv.ExistsUserByUsername(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return domain_errors.ErrExistsUsername
	}

	// Create the user
	err = s.user_serv.CreateUser(user)
	if err != nil {
		return err
	}

	return nil

}
