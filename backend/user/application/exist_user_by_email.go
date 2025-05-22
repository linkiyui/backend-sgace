package application

func (s *UserService) ExistsUserByEmail(email string) (bool, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}
	return true, nil
}
