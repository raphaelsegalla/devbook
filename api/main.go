package main

import (
	"api/src/config"
	"api/src/router"

	// "crypto/rand"
	// "encoding/base64"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	chave := make([]byte, 64)

// 	if _, erro := rand.Read(chave); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)

// 	fmt.Println(stringBase64)
// }

func main() {
	config.Carregar()

	// fmt.Println(config.Porta)
	// fmt.Println(config.StringConexaoBanco)

	fmt.Printf("Escutando na porta %d", config.Porta)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
	// log.Fatal(http.ListenAndServe(":5000", r))
}
