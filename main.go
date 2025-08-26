package main

import (
	"devbook-app/internal/config"
	"devbook-app/internal/cookies"
	"devbook-app/pkg/router"
	"devbook-app/pkg/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	app := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), app))
}
