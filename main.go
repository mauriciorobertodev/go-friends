package main

import (
	"fmt"
	"go-friends/pkg/config"
	"go-friends/pkg/router"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Rodando API...")

	config.Load()

	r := router.NewRouter()

	fmt.Printf("Escutando na porta: %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
