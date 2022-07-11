package server

import (
	"fmt"
	"net/http"
)

var tokenShit = "https://graph.facebook.com/oauth/access_token" +
	"?client_id={998747317492266}" +
	"&client_secret={a87ce8356d8788bac24d22ddc93442d5}" +
	"&grant_type"

var toFacebook = "https://graph.facebook.com/" +
	"v14.0/" + 
	// "oauth/access_token?grant_type={user_access_token}" +
	"oauth/access_token?grant_type={fb_exchange_token}" +
	"&client_id={998747317492266}" +
	"&client_secret={a87ce8356d8788bac24d22ddc93442d5}" +
	"&fb_exchange_token={EAAOMWvsUWioBAEZBzXGe3jYadilwT0VDZCSWsZC8CIcBAgRI7IGLopyqNzdf7XGpo3yJmg7PRrfBOZCPC5lnbMA1MRmT8gNZCJQAxQNmRgXAqi1HWY0lpujRpfzXsTU3VnxQ6zOH5YfkLZCRhbRhcDXX9ZCPIYyGdnBuDL5ValDnoQ4hfDIq3TyJVvWVJpK3rvObsGM0otgXgZDZD}"

func (s *Server) handleSendLLAT() http.HandlerFunc {
	fmt.Println("\nSending LLAT endpoint RUNNING")
	// var received_updates []int

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("\nSend LLAT endpoint HIT\n")
		defer s.handleUnexpectedError(w, r)

		fmt.Println("url: ", toFacebook)
		resp, err := http.Get(toFacebook)
		if err != nil {
			fmt.Println("Error sending LLAT get request: ", err)
			fmt.Println("response if any: ", resp)
			return
		}

		fmt.Println("no error sending request")
		fmt.Println("response: ", resp)
		
		// resp, err := http.Get("https:/")
	}
		
}