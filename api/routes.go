package api

func (s *Server) SetUpRoutes() {
	s.router.GET("/ping", s.Ping)
}
