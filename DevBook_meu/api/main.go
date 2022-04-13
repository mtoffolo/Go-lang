package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*

Gerado de Secret Key
func init() {

	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {

		log.Fatal(erro)

	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)

	fmt.Println(stringBase64)

}
*/
func main() {

	fmt.Println("Iniciando a API")
	config.Carregar()
	r := router.Gerar()
	fmt.Println("Escutando na Porta", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
