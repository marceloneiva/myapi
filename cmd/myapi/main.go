package main

import (
	"log"
	"net/http"

	"github.com/marceloneiva/myapi/internal/aplication/usecase"
	"github.com/marceloneiva/myapi/internal/config"
	"github.com/marceloneiva/myapi/internal/infrastructure/api"
	"github.com/marceloneiva/myapi/internal/infrastructure/repository"
)

func main() {
	config.Connect() // inicia a conexão com o banco

	repo := repository.NewMySQLRateRepo()
	uc := usecase.NewConvertCurrencyUseCase(repo)
	handler := api.NewHandler(uc)

	http.HandleFunc("/convert", handler.ConvertCurrency)

	log.Println("Servidor iniciado em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
