package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConexaoBanco é a string de conexão do banco postgres
	StringConexaoBanco = ""
	// Porta onde a API vai estar Rodando
	Porta = 0

	/*SecretKey é achave que será usada para assinar o token*/
	SecretKey []byte
)

/*Carregar vai iniciar as variaveis de ambiente*/
func Carregar() {
	var erro error

	if godotenv.Load(); erro != nil {

		log.Fatal(erro) /*Mata a execução*/

	}
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {

		Porta = 9000

	}

	StringConexaoBanco = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
