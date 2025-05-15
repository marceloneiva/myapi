package main

import (
    "log"
    "github.com/marceloneiva/myapi/internal/api"
)

func main() {
    //config.Connect() // inicia a conexão com o banco
    
    router := api.SetupRoutes()

    log.Println("Servidor iniciado na porta :8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Erro ao iniciar o servidor: %v", err)
    }
}