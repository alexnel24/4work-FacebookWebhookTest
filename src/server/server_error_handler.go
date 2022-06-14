package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ServerErrorHandler struct {
	err error
	status int
}

// HTTP error statuses with message
var(
	InternalServerError = &ServerErrorHandler{status: http.StatusInternalServerError, err: errors.New("SERVER ERROR: Internal Server Error. Try again later")}
	UnauthorizedIpError = &ServerErrorHandler{status: http.StatusForbidden, err: errors.New("APP ERROR: ip address not authorized to access this resource")}
	BadContextStoreRequestError = &ServerErrorHandler{status: http.StatusForbidden, err: errors.New("BAD REQUEST: email address expected in POST request params")}
	WantedPostRequestErr   = &ServerErrorHandler{status: http.StatusMethodNotAllowed, err: errors.New("POST REQUEST EXPECTED, OTHER TYPE RECEIVED")}
	OtherError = &ServerErrorHandler{status: http.StatusInternalServerError, err: errors.New("ERROR DECODING OBJECT")}
)

func (e *ServerErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.status)

	json.NewEncoder(w).Encode(map[string]string{"message": e.err.Error()})
}