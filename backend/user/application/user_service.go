package application

import (
	user_domain "github.com/sgace/backend/user/domain"
)

type IUserService interface {
	// CreateUser creates a new user
	CreateUser(user *user_domain.User) error
	ExistsUserByEmail(email string) (bool, error)
	ExistsUserByUsername(username string) (bool, error)
	GetUserByUsername(username string) (*user_domain.User, error)
	GetMyUserInfo(id string) (*user_domain.User, error)
	DeleteMyUser(id string) error
	UpdateMyUser(id string, user *user_domain.User) error
	GetUserByID(id string) (*user_domain.User, error)
	// GetUserProgress(id string) (*user_domain.UserProgress, error)
}

type UserService struct {
	userRepo user_domain.IUserRepository
}

func NewUserService(userRepo user_domain.IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
