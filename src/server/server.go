package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server{
	s := &Server{
		http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	s.router.ServeHTTP(w, r)
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int){
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if data != nil{
		if err := json.NewEncoder(w).Encode(data); err != nil {
			InternalServerError.ServeHTTP(w, r)
			fmt.Println("Internal Server Error Occured: ", err)
		}
	}
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		fmt.Println("body of request is nil:", r)
		return errors.New("body of request is nil")
	}
	decoder := json.NewDecoder(r.Body)
	fmt.Println("decoder: ", decoder)
	decoded := decoder.Decode(v)
	fmt.Println("decoded: ", decoded)
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *Server) handleUnexpectedError(w http.ResponseWriter, r *http.Request){
	if recover_err := recover(); recover_err != nil {
		InternalServerError.ServeHTTP(w, r)
		fmt.Println("Internal Server Error Occured: ", recover_err)
	}
}