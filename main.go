package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	port := getPort()
	// setGetStartedPayload("GET_STARTED")
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/webhook", webhookGetHandler).Methods("GET")
	router.HandleFunc("/webhook", webhookPostHandler).Methods("POST")
	fmt.Printf("Server up and running. Running on PORT: %s\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("error listening to server: ", err)
	}
}
func indexHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request)
	fmt.Fprint(response, "Got my server up and running in Go. Yay!!")
}
func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		/**
		*TODO: get the port declared in the yml config.
		 */
		port = "3500"
		fmt.Printf("PORT NOT DEFINED. USING THE PORT %s as the running port\n", port)
	}
	return ":" + port
}
