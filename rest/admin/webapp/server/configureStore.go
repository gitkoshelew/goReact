package server

import "admin/domain/store"

func (s *Server) configureStore() error {
	st := store.New(s.config)
	if err := st.Open(); err != nil {
		st.Logger.Errorf("Error occurred while configuring store. Err msg:%v.", err)
		return err
	}

	s.Store = st

	return nil
}
