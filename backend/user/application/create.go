package application

import (
	"time"

	user_domain "github.com/sgace/backend/user/domain"
	"github.com/sgace/utils"
)

func (s *UserService) CreateUser(user *user_domain.User) error {

	hashedPassword := utils.HashingPasswordFunc(user.Password)

	user.Password = hashedPassword

	id, err := utils.GenerateUUIDv7()
	if err != nil {
		return err
	}

	user_to_postgres := &user_domain.User{
		ID:        id,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
	}

	// Create the user in the repository
	if err := s.userRepo.CreateUser(user_to_postgres); err != nil {
		return err
	}

	return nil
}
