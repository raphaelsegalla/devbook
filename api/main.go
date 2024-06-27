package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	// fmt.Println(config.Porta)
	// fmt.Println(config.StringConexaoBanco)

	fmt.Printf("Escutando na porta %d", config.Porta)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
	// log.Fatal(http.ListenAndServe(":5000", r))
}
