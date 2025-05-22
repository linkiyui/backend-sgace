package application

func (s *SolicitudeService) GetAcceptedSolicitudes(id string) (int64, error) {
	sol, err := s.SolicitudeRepo.GetAcceptedSolicitudes(id)

	if err != nil {
		return 0, err
	}

	return int64(len(sol)), nil

}
