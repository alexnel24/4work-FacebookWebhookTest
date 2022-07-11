package server



func (s Server) routes(){
	s.router.HandleFunc("/health-check", s.handleHealthCheck())
	s.router.HandleFunc("/facebook", s.handleFacebook())
	// s.router.HandleFunc("/send-llat", s.handleSendLLAT())

	// s.router.HandleFunc("/get-long-access-token", s.handleGetLongAccessToken())

}