package application

import (
	user_domain "github.com/sgace/backend/user/domain"
	domain_errors "github.com/sgace/errors"
)

func (s *UserService) GetMyUserInfo(id string) (*user_domain.User, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain_errors.ErrNotFound
	}
	return user, nil
}
