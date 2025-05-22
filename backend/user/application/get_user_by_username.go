package application

import (
	user_domain "github.com/sgace/backend/user/domain"
)

func (s *UserService) GetUserByUsername(username string) (*user_domain.User, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}
