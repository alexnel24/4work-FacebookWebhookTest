package server

import (
	"fmt"
	"net/http"
	"reflect"
)

func (s *Server) handleFacebook() http.HandlerFunc {
	fmt.Println("\nFacebook endpoint RUNNING")

	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method{

		
		
		case http.MethodGet:
			fmt.Println("\nFacebook endpoint HIT w GET request\n")
			fmt.Println("Request: ", r)
			defer s.handleUnexpectedError(w, r)

			test_value := r.URL.Query()
			fmt.Println("test_query value: ", test_value)
			fmt.Println("test_query type: ", reflect.TypeOf(test_value))

		
		case http.MethodPost:
			fmt.Println("\nFacebook endpoint HIT w POST request")
			fmt.Println("Request: ", r)
			defer s.handleUnexpectedError(w, r)



		}
		
		_, err := w.Write([]byte("FACEBOOK ENDPOINT OK"))
		if err != nil {
			InternalServerError.ServeHTTP(w, r)
		}
	}
}