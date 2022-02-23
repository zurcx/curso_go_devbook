package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (

	// APIURL representa a url para comunicaçao com a API
	APIURL = ""
	// Porta onde a Api esta rodando
	Porta = 0
	// HashKey utilizada para autenticar o cookie
	HashKey []byte
	// BlockKey utilizada para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializa as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
