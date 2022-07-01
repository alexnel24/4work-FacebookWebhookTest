package server

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	// "io"
	"net/http"
	"reflect"
)

type facebookRequest struct {
	Object string `json:"object"`
	Entry []ActualHeaders `json:"entry"`
}

// type FacebookFirstLead struct {
// 	ZeroOnly []ActualHeaders `json:"0"`
// 	// OneOnly []ActualLead `json:"1"`
// }

type ActualHeaders struct{
	Id string `json:"id"`
	Time int `json:"time"`
	Uid string `json:"uid"`
	ActualLead []ActualLead `json:"changes"`
}
type ActualLead struct {
	Field string `json:"field"`
	Value ActualLeadValues `json:"value"`
}


type ActualLeadValues struct {
	Adid string `json:"ad_id"`
	Form_id string `json:"form_id"`
	Leadgen string `json:"leadgen_id"`
	Created_time int `json:"created_time"`
	Page_id string `json:"page_id"`
	Adgroup_id string `json:"adgroup_id"`
}


func (s *Server) handleFacebook() http.HandlerFunc {
	fmt.Println("\nFacebook endpoint RUNNING")
	// var received_updates []int

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("http request received")

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


			var step1 facebookRequest
			json.NewDecoder(r.Body).Decode(&step1)

			fmt.Println(step1.Entry[0].Id)
			fmt.Println(step1.Entry[0].Time)
			fmt.Println(step1.Entry[0].Uid)
			fmt.Println(step1.Entry[0].ActualLead[0])
			fmt.Println(step1.Entry[0].ActualLead[0].Field)
			fmt.Println(step1.Entry[0].ActualLead[0].Value)
			fmt.Println("holy")

			fmt.Println(step1.Entry[0].ActualLead[0].Value.Adid)
			fmt.Println(step1.Entry[0].ActualLead[0].Value.Form_id)
			fmt.Println(step1.Entry[0].ActualLead[0].Value.Leadgen)
			fmt.Println(step1.Object)
			
			
			// var step2 FacebookFirstLead
			// json.NewDecoder(r.Body).Decode(&step2)
			
			// fmt.Println(step2.ZeroOnly)

			// var eb ActualHeaders
			// json.NewDecoder(r.Body).Decode(&eb)
	
			// fmt.Println(eb.Id)
			// fmt.Println(eb.AnotherKey)
	
			// fmt.Fprintf(w, "%s %s came from post body", eb.Key, eb.AnotherKey)
	


			// new_byte_arr, err := io.ReadAll(r.Body)
			// if err != nil {
			// 	fmt.Println("readAll err: ", err)
			// 	// OtherError.ServeHTTP(w, r)
			// 	// return
			// }
			// fmt.Println("IO READALL: ", new_byte_arr)
			// fmt.Println("readall type: ", reflect.TypeOf(new_byte_arr))

			// wr := WebHookRequest{}
			// var objs []map[string]*json.RawMessage



			// var objs *json.RawMessage
			// err = json.Unmarshal(new_byte_arr, &objs)

			



			// err = s.decode(w, r, &incoming)
			// if err != nil {
			// 	// fmt.Println("decode err: ", err)
			// 	fmt.Println("incoming object: ", objs)
			// 	OtherError.ServeHTTP(w, r)
			// 	return
			// }

			fmt.Println("items successfully received")
			// fmt.Println("incoming object: ", eb.Id)





		}
		
		// _, err := w.Write([]byte("FACEBOOK ENDPOINT OK"))
		// if err != nil {
		// 	InternalServerError.ServeHTTP(w, r)
		// }
	}
}