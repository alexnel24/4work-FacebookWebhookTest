package server

import (
	"fmt"
	"net/http"
)

func (s *Server) handleHealthCheck() http.HandlerFunc {
	fmt.Println("\nhealth-check endpoint RUNNING")

	return func(w http.ResponseWriter, r *http.Request) {
		
		fmt.Println("\nhealth-check hit")
		_, err := w.Write([]byte("HEALTH CHECK OK"))
		if err != nil {
			InternalServerError.ServeHTTP(w, r)
		}
	}
}