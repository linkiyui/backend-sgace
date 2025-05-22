package errors

var errors = map[string]*DomainError{}

type DomainError struct {
	Code        int
	Message     string
	Description string
}

func (de DomainError) Error() string {
	return de.Message
}

func IsDomainError(e any) *DomainError {
	if de, ok := e.(DomainError); ok {
		return &de
	}
	if de, ok := e.(*DomainError); ok {
		return de
	}
	return nil
}

func (de DomainError) Equal(e *DomainError) bool {
	return de.Message == e.Message
}

func newDE(code int, message, description string) DomainError {
	de := &DomainError{
		Code:        code,
		Message:     message,
		Description: description,
	}
	errors[message] = de
	return *de
}

var (
	ErrUserNotFound       = newDE(404, "User not found", "User not found")
	ErrExistsUsername     = newDE(409, "Username already exists", "Username already exists")
	ErrNotFound           = newDE(404, "Not found", "Not found")
	ErrInvalidPassword    = newDE(401, "Invalid password", "Invalid password")
	ErrEmailAlreadyExists = newDE(409, "Email already exists", "Email already exists")
	ErrbadRequest         = newDE(400, "Bad request", "Bad request")
)
