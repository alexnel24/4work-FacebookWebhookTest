package server

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

type FacebookRequest struct {
	Object string `json:"object"`
	Entry []FacebookEntry `json:"entry"`
}

type FacebookEntry struct {
	Id string `json:"id"`
	Uid string `json:"uid"`
	Changes []string `json:"changed_fields"`
	
}

// type FacebookField struct{
// 	Field string `json:"field"`
// 	Value []FacebookFieldValue `json:"value"`
// }

// type FacebookFieldValue struct {
// 	Ad_id string `json:"ad_id"`
// 	Form_id string `json:"formt_id"`
// 	Leadgen_id string `json:"leadgen_id"`
// 	Created_Time string ``
// }








func (s *Server) handleFacebook() http.HandlerFunc {
	fmt.Println("\nFacebook endpoint RUNNING")
	// var received_updates []int

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
			// fmt.Println("Request: ", r)
			// fmt.Println("Request body: ", r.Body)
			// fmt.Println("request body type: ", reflect.TypeOf(r.Body))
			defer s.handleUnexpectedError(w, r)
			// received_updates.unshift(r.Body)

			// var incoming facebookRequest


			new_byte_arr, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("readAll err: ", err)
				// OtherError.ServeHTTP(w, r)
				// return
			}
			// fmt.Println("IO READALL: ", new_byte_arr)
			// fmt.Println("readall type: ", reflect.TypeOf(new_byte_arr))

			// wr := WebHookRequest{}
			// var objs []map[string]*json.RawMessage
			var objs FacebookRequest
			err = json.Unmarshal(new_byte_arr, &objs)

			



			// err = s.decode(w, r, &incoming)
			if err != nil {
				fmt.Println("decode err: ", err)
				fmt.Println("incoming object: ", objs)
				OtherError.ServeHTTP(w, r)
				return
			}

			fmt.Println("items successfully received")
			fmt.Println("incoming object: ", objs)


		}
		
		// _, err := w.Write([]byte("FACEBOOK ENDPOINT OK"))
		// if err != nil {
		// 	InternalServerError.ServeHTTP(w, r)
		// }
	}
}