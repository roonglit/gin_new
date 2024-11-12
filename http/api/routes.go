package api

func (s *Server) SetUpRoutes() {
	s.router.GET("/users/greet", s.GreetUser)
}
