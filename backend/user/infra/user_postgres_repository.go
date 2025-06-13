package repository

import (
	"fmt"
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

	fmt.Println(user)

	result := r.db.Table("users").Create(map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"role":       user.Role,
		"email":      user.Email,
		"password":   user.Password,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		// Agrega aquí todos los campos de tu estructura User
	})

	if result.Error != nil {
		return result.Error
	}

	// Si necesitas el ID generado (para auto-increment)
	if result.RowsAffected == 1 && user.ID == "" {
		r.db.Table("users").Select("id").Last(&user)
	}

	return nil
}

func (r *UserPostgresRepository) GetUserByEmail(email string) (*user_domain.User, error) {
	var user user_domain.User
	if err := r.db.Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgresRepository) GetUserByUsername(username string) (*user_domain.User, error) {
	var user user_domain.User
	if err := r.db.Table("users").Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgresRepository) GetUserByID(id string) (*user_domain.User, error) {
	var user user_domain.User
	if err := r.db.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgresRepository) DeleteMyUser(id string) error {
	var user user_domain.User
	if err := r.db.Table("users").Where("id = ?", id).Delete(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}

func (r *UserPostgresRepository) UpdateMyUser(id string, user *user_domain.User) error {
	var userToUpdate user_domain.User
	if err := r.db.Table("users").Where("id = ?", id).First(&userToUpdate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	userToUpdate.Username = user.Username
	userToUpdate.Email = user.Email
	userToUpdate.Password = user.Password
	userToUpdate.UpdatedAt = time.Now()
	if err := r.db.Table("users").Where("id = ?", id).Save(&userToUpdate).Error; err != nil {
		return err
	}
	return nil
}
