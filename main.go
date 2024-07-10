package main

import (
	"go-assessment/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/files", handlers.UploadHandler)
	http.HandleFunc("/v1/files/get", handlers.GetCIDHandler)
	log.Println("Server listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
