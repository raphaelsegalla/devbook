package main

import (
	"api/src/config"
	"api/src/router"

	// "crypto/rand"
	// "encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
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

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.NewRelicAppName),
		newrelic.ConfigLicense(config.NewRelicLicense),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, New Relic!")
	}))

	http.HandleFunc("/publicacoes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Começar uma nova transação New Relic
			txn := app.StartTransaction("POST /publicacoes")
			defer txn.End()

			// Aqui vai o processamento da sua requisição
			fmt.Fprintf(w, "Requisição POST recebida!")
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Printf("Escutando na porta %d \n", config.Porta)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
	// log.Fatal(http.ListenAndServe(":5000", r))
}
