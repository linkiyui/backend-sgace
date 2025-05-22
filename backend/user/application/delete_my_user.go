package application

func (s *UserService) DeleteMyUser(id string) error {
	err := s.userRepo.DeleteMyUser(id)
	if err != nil {
		return err
	}
	return nil
}
