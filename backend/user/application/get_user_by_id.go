package application

import (
	user_domain "github.com/sgace/backend/user/domain"
)

func (s *UserService) GetUserByID(id string) (*user_domain.User, error) {
	return s.userRepo.GetUserByID(id)
}
