package server



func (s Server) routes(){
	s.router.HandleFunc("/health-check", s.handleHealthCheck())

}