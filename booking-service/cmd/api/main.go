package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello World from booking")

	port := os.Getenv("BOOKING_SERVICE_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Running server on port " + port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), http.FileServer(http.Dir("public")))
	if err != nil {
		return
	}
}
