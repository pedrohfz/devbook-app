package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL representa a URL para comunicação com a API.
	APIURL = ""

	// Porta representa a porta onde a aplicação Web está rodando.
	Porta = 0

	// HashKey é utilizada para autenticar o cookie.
	HashKey []byte

	// BlockKey é utilizada para criptografar os dados do cookie.
	BlockKey []byte
)

// Carregar() inicializa as variáveis de ambiente.
func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("HASH_KEY"))
}
