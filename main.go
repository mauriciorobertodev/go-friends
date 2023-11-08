package main

import (
	"fmt"
	"go-friends/pkg/router"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Iniciando API...")
	
	r := router.NewRouter()

	err := http.ListenAndServe(":5000", r)

	if err != nil {
		log.Fatal("Erro ao iniciar o servidor http")
	}

	fmt.Printf("API escutando na porta 5000...")
}