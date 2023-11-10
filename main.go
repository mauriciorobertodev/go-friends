package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"go-friends/pkg/config"
	"go-friends/pkg/router"
	"log"
	"net/http"
)

func init() {
	chave := make([]byte, 64)

	if _, err := rand.Read(chave); err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(chave))
}

func main() {
	fmt.Printf("Rodando API...")

	config.Load()

	r := router.NewRouter()

	fmt.Printf("Escutando na porta: %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
