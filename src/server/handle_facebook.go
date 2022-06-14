package server

import (
	"fmt"
	"net/http"
	"reflect"
)

func (s *Server) handleFacebook() http.HandlerFunc {
	fmt.Println("\nFacebook endpoint RUNNING")
	var received_updates []int

	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method{

		
		
		case http.MethodGet:
			fmt.Println("\nFacebook endpoint HIT w GET request\n")
			fmt.Println("Request: ", r)
			defer s.handleUnexpectedError(w, r)

			test_value := r.URL.Query()
			fmt.Println("test_query value: ", test_value)
			fmt.Println("test_query type: ", reflect.TypeOf(test_value))

			test_url_hubMode := test_value.Get("hub.mode")
			fmt.Println("hub.mode: ", test_url_hubMode)

			test_url_token := test_value.Get("hub.verify_token")
			fmt.Println("hub.verify_token: ", test_url_token)

			test_url_challenge := test_value.Get("hub.challenge")
			fmt.Println("hub.challenge: ", test_url_challenge)
			w.Write([]byte(test_url_challenge))
			

		
		case http.MethodPost:
			fmt.Println("\nFacebook endpoint HIT w POST request")
			fmt.Println("Request: ", r)
			fmt.Println("Request body: ", r.Body)
			defer s.handleUnexpectedError(w, r)
			// received_updates.unshift(r.Body)



		}
		
		// _, err := w.Write([]byte("FACEBOOK ENDPOINT OK"))
		// if err != nil {
		// 	InternalServerError.ServeHTTP(w, r)
		// }
	}
}