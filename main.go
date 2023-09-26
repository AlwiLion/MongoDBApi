package main

import (
	"fmt"
	"log"
	router "main/Router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB Connection")

	r := router.Router()
	fmt.Println("Starting Server")

	log.Fatal(http.ListenAndServe(":4001", r))
	fmt.Println("Listning on Port 4000")
}
