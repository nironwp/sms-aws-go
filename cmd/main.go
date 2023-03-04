package main

import (
	"fmt"
	"log"
	"net/http"
	"sms-aws-go/cmd/http/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	// Define suas rotas aqui usando router.HandleFunc() ou router.HandleFunc().Methods()
	router.HandleFunc("/sms", routes.SendSMSHandler).Methods("POST")
	fmt.Println("Listening in port :3030")
	log.Fatal(http.ListenAndServe(":3030", router))
}
