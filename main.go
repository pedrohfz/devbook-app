package main

import (
	"devbook-app/pkg/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	app := router.Gerar()

	fmt.Println("Rodando WebApp!")
	log.Fatal(http.ListenAndServe(":3000", app))
}
