package application

import (
	user_domain "github.com/sgace/backend/user/domain"
	domain_errors "github.com/sgace/errors"
	"github.com/sgace/utils"
)

func (s *UserService) UpdateMyUser(id string, user *user_domain.User) error {

	user_from_postgres, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}

	if user_from_postgres == nil {
		return domain_errors.ErrNotFound
	}

	hashed_password := user_from_postgres.Password

	if utils.CheckPasswordHashFunc(user.Password, hashed_password) {
		return domain_errors.ErrInvalidPassword
	}

	hashed_password = utils.HashingPasswordFunc(user.Password)

	user.Password = hashed_password

	return s.userRepo.UpdateMyUser(id, user)
}
