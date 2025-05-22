package application

func (s *UserService) ExistsUserByUsername(username string) (bool, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}
	return true, nil
}
