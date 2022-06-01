package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alexnel24/4work-FacebookWebhookTest/src/server"
)

func main(){
	fmt.Println("Hi Alex")
	port := os.Getenv("PORT")

	server := server.NewServer()

	log.Fatal(http.ListenAndServe(":"+port, server))
}