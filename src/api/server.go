package api

import (
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	fmt.Println("\n\tListening [::]:3000\n")
	r := router.New()
	log.Fatal(http.ListenAndServe(":3000", r))
}