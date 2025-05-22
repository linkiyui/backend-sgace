package repository

import (
	"time"

	user_domain "github.com/sgace/backend/user/domain"
	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{
		db: db,
	}
}

func (r *UserPostgresRepository) CreateUser(user *user_domain.User) error {
	if err := r.db.Table("user").Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserPostgresRepository) GetUserByEmail(email string) (*user_domain.User, error) {
	var user user_domain.User
	if err := r.db.Table("user").Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgresRepository) GetUserByUsername(username string) (*user_domain.User, error) {
	var user user_domain.User
	if err := r.db.Table("user").Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgresRepository) GetUserByID(id string) (*user_domain.User, error) {
	var user user_domain.User
	if err := r.db.Table("user").Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgresRepository) DeleteMyUser(id string) error {
	var user user_domain.User
	if err := r.db.Table("user").Where("id = ?", id).Delete(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}

func (r *UserPostgresRepository) UpdateMyUser(id string, user *user_domain.User) error {
	var userToUpdate user_domain.User
	if err := r.db.Table("user").Where("id = ?", id).First(&userToUpdate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	userToUpdate.Username = user.Username
	userToUpdate.Email = user.Email
	userToUpdate.Password = user.Password
	userToUpdate.UpdatedAt = time.Now()
	if err := r.db.Table("user").Where("id = ?", id).Save(&userToUpdate).Error; err != nil {
		return err
	}
	return nil
}
