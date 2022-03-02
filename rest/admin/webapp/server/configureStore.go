package server

import "admin/domain/store"

func (s *Server) configureStore() error {
	st := store.New(s.config)
	if err := st.Open(); err != nil {
		return err
	}

	s.Store = st

	return nil
}
