package domain

type IUserRepository interface {
	// CreateUser creates a new user in the repository
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id string) (*User, error)
	DeleteMyUser(id string) error
	UpdateMyUser(id string, user *User) error
}
